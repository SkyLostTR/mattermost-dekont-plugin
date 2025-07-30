// Package main implements a Mattermost plugin for parsing PDF bank receipts.
//
// PDF Dekont Parser Plugin for Mattermost
// Developed by SkyLostTR (@Keeftraum) with ❤️
//
// This plugin automatically extracts transaction details from PDF bank receipts
// and supports multiple Turkish banks including İş Bankası, Garanti BBVA,
// Akbank, Yapı Kredi, Ziraat Bankası, VakıfBank, Kuveyt Türk, and HalkBank.
//
// For support, feature requests, or to report issues:
// https://github.com/SkyLostTR/mattermost-dekont-plugin
//
// Author: SkyLostTR (@Keeftraum)
// License: See LICENSE file in the project root
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ledongthuc/pdf"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

// Configuration contains the plugin configuration settings
type Configuration struct {
	EnablePlugin             bool   `json:"EnablePlugin"`
	ProcessOnlyInChannels    bool   `json:"ProcessOnlyInChannels"`
	AllowedChannels          string `json:"AllowedChannels"`
	MaxFileSizeMB            int    `json:"MaxFileSizeMB"`
	CustomMessagePrefix      string `json:"CustomMessagePrefix"`
	IncludeTimestamp         bool   `json:"IncludeTimestamp"`
	NotifyOnProcessingError  bool   `json:"NotifyOnProcessingError"`
	ErrorNotificationMessage string `json:"ErrorNotificationMessage"`
	EnableDebugLogging       bool   `json:"EnableDebugLogging"`
	SupportedBanks           string `json:"SupportedBanks"`
}

// Plugin represents the main plugin instance.
// Developed by SkyLostTR (@Keeftraum) for the Mattermost community
type Plugin struct {
	plugin.MattermostPlugin
	configuration *Configuration
}

// OnActivate is called when the plugin is activated.
func (p *Plugin) OnActivate() error {
	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	p.API.LogInfo("PDF Dekont Parser Plugin activated successfully",
		"version", "1.0.0",
		"author", "SkyLostTR (@Keeftraum)",
		"repository", "https://github.com/SkyLostTR/mattermost-dekont-plugin")

	return nil
}

// OnConfigurationChange is called when the plugin configuration changes
func (p *Plugin) OnConfigurationChange() error {
	var configuration = new(Configuration)

	if err := p.API.LoadPluginConfiguration(configuration); err != nil {
		p.API.LogError("Failed to load plugin configuration", "error", err.Error())
		return err
	}

	// Set default values if not configured
	if configuration.MaxFileSizeMB == 0 {
		configuration.MaxFileSizeMB = 10
	}
	if configuration.CustomMessagePrefix == "" {
		configuration.CustomMessagePrefix = "📄 **Dekont Bilgileri:**"
	}
	if configuration.ErrorNotificationMessage == "" {
		configuration.ErrorNotificationMessage = "⚠️ PDF dekont işlenirken hata oluştu. Lütfen dosyanın geçerli bir banka dekontu olduğundan emin olun."
	}

	p.configuration = configuration

	if configuration.EnableDebugLogging {
		p.API.LogDebug("Plugin configuration updated",
			"EnablePlugin", configuration.EnablePlugin,
			"ProcessOnlyInChannels", configuration.ProcessOnlyInChannels,
			"MaxFileSizeMB", configuration.MaxFileSizeMB)
	}

	return nil
}

// getConfiguration retrieves the active configuration under lock
func (p *Plugin) getConfiguration() *Configuration {
	if p.configuration == nil {
		return &Configuration{}
	}
	return p.configuration
}

// MessageHasBeenPosted processes newly posted messages to extract PDF content.
func (p *Plugin) MessageHasBeenPosted(_ *plugin.Context, post *model.Post) {
	config := p.getConfiguration()

	// Check if plugin is enabled
	if !config.EnablePlugin {
		return
	}

	if post.Type != "" || post.FileIds == nil || len(post.FileIds) == 0 {
		return
	}

	// Check if we should only process in specific channels
	if config.ProcessOnlyInChannels && config.AllowedChannels != "" {
		channel, err := p.API.GetChannel(post.ChannelId)
		if err != nil {
			p.API.LogError("Failed to get channel info", "channelId", post.ChannelId, "error", err.Error())
			return
		}

		allowedChannels := strings.Split(config.AllowedChannels, ",")
		isAllowed := false
		for _, allowedChannel := range allowedChannels {
			if strings.TrimSpace(allowedChannel) == channel.Name {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			if config.EnableDebugLogging {
				p.API.LogDebug("Skipping PDF processing - channel not in allowed list",
					"channel", channel.Name,
					"allowedChannels", config.AllowedChannels)
			}
			return
		}
	}

	for _, fileID := range post.FileIds {
		if err := p.processFileUpload(fileID, post); err != nil {
			p.API.LogError("Failed to process file upload",
				"fileId", fileID,
				"error", err.Error())

			// Send error notification if enabled
			if config.NotifyOnProcessingError {
				p.sendErrorNotification(post.ChannelId, config.ErrorNotificationMessage)
			}
		}
	}
}

// sendErrorNotification sends an error message to the channel
func (p *Plugin) sendErrorNotification(channelID, message string) {
	post := &model.Post{
		ChannelId: channelID,
		Message:   message,
		Type:      "custom_pdf_error",
	}

	if _, err := p.API.CreatePost(post); err != nil {
		p.API.LogError("Failed to send error notification", "error", err.Error())
	}
}

func (p *Plugin) processFileUpload(fileID string, post *model.Post) error {
	config := p.getConfiguration()

	fileInfo, err := p.API.GetFileInfo(fileID)
	if err != nil || !strings.HasSuffix(fileInfo.Name, ".pdf") {
		return nil // Not a PDF file, skip silently
	}

	if config.EnableDebugLogging {
		p.API.LogDebug("Processing PDF file",
			"fileName", fileInfo.Name,
			"fileSize", fileInfo.Size,
			"author", "SkyLostTR (@Keeftraum)")
	}

	// Check file size limit
	maxSizeBytes := int64(config.MaxFileSizeMB * 1024 * 1024)
	if fileInfo.Size > maxSizeBytes {
		if config.EnableDebugLogging {
			p.API.LogDebug("Skipping PDF - file too large",
				"fileName", fileInfo.Name,
				"fileSize", fileInfo.Size,
				"maxSize", maxSizeBytes)
		}
		return nil
	}

	data, appErr := p.API.GetFile(fileID)
	if appErr != nil {
		return appErr
	}

	tempFile, fileErr := os.CreateTemp("", "*.pdf")
	if fileErr != nil {
		return fileErr
	}
	defer func() {
		if err := tempFile.Close(); err != nil {
			p.API.LogError("Failed to close temp file", "error", err.Error())
		}
		if err := os.Remove(tempFile.Name()); err != nil {
			p.API.LogError("Failed to remove temp file", "error", err.Error())
		}
	}()

	if _, writeErr := tempFile.Write(data); writeErr != nil {
		return writeErr
	}

	if closeErr := tempFile.Close(); closeErr != nil {
		return closeErr
	}

	file, r, pdfErr := pdf.Open(tempFile.Name())
	if pdfErr != nil {
		return pdfErr
	}
	defer file.Close()

	var extractedText string
	if r.NumPage() > 0 {
		page := r.Page(1)
		if page.V.IsNull() {
			return nil
		}
		var textErr error
		extractedText, textErr = page.GetPlainText(nil)
		if textErr != nil {
			return textErr
		}
	}

	description := p.extractFields(extractedText, config)
	if description != "" {
		// Add custom prefix and timestamp if enabled
		var fullMessage strings.Builder

		if config.CustomMessagePrefix != "" {
			fullMessage.WriteString(config.CustomMessagePrefix)
			fullMessage.WriteString("\n\n")
		}

		fullMessage.WriteString(description)

		if config.IncludeTimestamp {
			timestamp := time.Now().Format("02.01.2006 15:04:05")
			fullMessage.WriteString(fmt.Sprintf("\n\n*İşlenme Zamanı: %s*", timestamp))
		}

		// Add credits footer
		fullMessage.WriteString("\n\n---\n*Mattermost PDF Parser Plugin by SkyLostTR* 🚀")

		post.Message = fullMessage.String()
		_, appErr = p.API.UpdatePost(post)
		if appErr != nil {
			return appErr
		}

		if config.EnableDebugLogging {
			p.API.LogDebug("Successfully processed PDF and updated post",
				"fileName", fileInfo.Name,
				"extractedFields", len(strings.Split(description, "\n")))
		}
	}

	return nil
}

// extractFields extracts transaction details from PDF text
// Enhanced by SkyLostTR (@Keeftraum) to support multiple Turkish bank formats
func (p *Plugin) extractFields(text string, config *Configuration) string {
	var alici, gonderen, aciklama, tutar, tarih string

	if config.EnableDebugLogging {
		p.API.LogDebug("Starting field extraction",
			"textLength", len(text),
			"author", "SkyLostTR (@Keeftraum)")
	}

	// Enhanced regex patterns for multiple bank formats
	// VakıfBank patterns
	reAliciVakif := regexp.MustCompile(`(?i)ALICI\s*(?:AD\s*SOYAD/UNVAN)?\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reGonderenVakif := regexp.MustCompile(`(?i)G[ÖO]NDEREN\s*(?:AD\s*SOYAD\s*/?\s*UNVAN)?\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reAciklamaVakif := regexp.MustCompile(`(?i)İ[ŞS]LEM\s*A[ÇC]IKLAMASI\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reTutarVakif := regexp.MustCompile(`(?i)İ[ŞS]LEM\s*TUTARI\s*[:\-]?\s*(?:.*?)?([0-9]+(?:[.,][0-9]{1,3})*(?:[.,][0-9]{2})?)\s*(?:TL|₺)?`)
	reTarihVakif := regexp.MustCompile(`(?i)İ[ŞS]LEM\s*TARİHİ\s*[:\-]?\s*(.+?)(?:\n|$)`)

	// YapıKredi patterns
	reAliciYapi := regexp.MustCompile(`(?i)ALICI\s*ADI\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reGonderenYapi := regexp.MustCompile(`(?i)G[ÖO]NDEREN\s*ADI\s*SOYAD\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reAciklamaYapi := regexp.MustCompile(`(?i)A[ÇC]IKLAMA\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reTutarYapi := regexp.MustCompile(`(?i)G[İI]DEN\s*EFT\s*TUTARI\s*[:\-]?\s*(?:.*?)?([0-9]+(?:[.,][0-9]{1,3})*(?:[.,][0-9]{2})?)\s*(?:TL|₺)?`)

	// Kuveyt Türk patterns
	reAliciKuveyt := regexp.MustCompile(`(?i)ALICI\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reGonderenKuveyt := regexp.MustCompile(`(?i)G[ÖO]NDEREN\s*(?:KİŞİ|KISI)\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reAciklamaKuveyt := regexp.MustCompile(`(?i)A[ÇC]IKLAMA\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reTutarKuveyt := regexp.MustCompile(`(?i)TUTAR\s*[:\-]?\s*(?:.*?)?([0-9]+(?:[.,][0-9]{1,3})*(?:[.,][0-9]{2})?)\s*(?:TL|₺)?`)

	// HalkBank patterns
	reAliciHalk := regexp.MustCompile(`(?i)ALICI\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reGonderenHalk := regexp.MustCompile(`(?i)G[ÖO]NDEREN\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reAciklamaHalk := regexp.MustCompile(`(?i)A[ÇC]IKLAMA\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reTutarHalk := regexp.MustCompile(`(?i)İ[ŞS]LEM\s*TUTARI\s*\(TL\)\s*[:\-]?\s*(?:.*?)?([0-9]+(?:[.,][0-9]{1,3})*(?:[.,][0-9]{2})?)\s*(?:TL|₺)?`)
	reTarihHalk := regexp.MustCompile(`(?i)İ[ŞS]LEM\s*TARİHİ\s*[:\-]?\s*(.+?)(?:\n|$)`)

	// Generic patterns (for existing banks and fallback)
	reAlici := regexp.MustCompile(`(?i)ALICI\s*(?:AD\s*SOYAD/UNVAN)?\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reAciklama := regexp.MustCompile(`(?i)A[CÇ]IKLAMA\s*[:\-]?\s*(.+?)(?:\n|$)`)
	reTutar := regexp.MustCompile(`(?i)(?:İ[ŞS]LEM\s*TUTARI|I[ŞS]LEM\s*TUTARI|TUTAR[IİĞ]?|HAVALE\s*TUTARI|G[İI]DEN\s*EFT\s*TUTARI|EFT\s*TUTARI|TRANSFER\s*TUTARI|PARA\s*TUTARI|M[İI]KTAR)\s*(?:\(TL\))?\s*[:\-\s]*(?:.*?)?([0-9]+(?:[.,][0-9]{1,3})*(?:[.,][0-9]{2})?)\s*(?:TL|₺)?`)

	// Try bank-specific patterns first, then fall back to generic patterns

	// VakıfBank
	if m := reAliciVakif.FindStringSubmatch(text); len(m) > 1 {
		alici = strings.TrimSpace(m[1])
	}
	if m := reGonderenVakif.FindStringSubmatch(text); len(m) > 1 {
		gonderen = strings.TrimSpace(m[1])
	}
	if m := reAciklamaVakif.FindStringSubmatch(text); len(m) > 1 {
		aciklama = strings.TrimSpace(m[1])
	}
	if m := reTutarVakif.FindStringSubmatch(text); len(m) > 1 {
		tutar = strings.TrimSpace(m[1])
	}
	if m := reTarihVakif.FindStringSubmatch(text); len(m) > 1 {
		tarih = strings.TrimSpace(m[1])
	}

	// YapıKredi
	if alici == "" {
		if m := reAliciYapi.FindStringSubmatch(text); len(m) > 1 {
			alici = strings.TrimSpace(m[1])
		}
	}
	if gonderen == "" {
		if m := reGonderenYapi.FindStringSubmatch(text); len(m) > 1 {
			gonderen = strings.TrimSpace(m[1])
		}
	}
	if aciklama == "" {
		if m := reAciklamaYapi.FindStringSubmatch(text); len(m) > 1 {
			aciklama = strings.TrimSpace(m[1])
		}
	}
	if tutar == "" {
		if m := reTutarYapi.FindStringSubmatch(text); len(m) > 1 {
			tutar = strings.TrimSpace(m[1])
		}
	}

	// Kuveyt Türk
	if alici == "" {
		if m := reAliciKuveyt.FindStringSubmatch(text); len(m) > 1 {
			alici = strings.TrimSpace(m[1])
		}
	}
	if gonderen == "" {
		if m := reGonderenKuveyt.FindStringSubmatch(text); len(m) > 1 {
			gonderen = strings.TrimSpace(m[1])
		}
	}
	if aciklama == "" {
		if m := reAciklamaKuveyt.FindStringSubmatch(text); len(m) > 1 {
			aciklama = strings.TrimSpace(m[1])
		}
	}
	if tutar == "" {
		if m := reTutarKuveyt.FindStringSubmatch(text); len(m) > 1 {
			tutar = strings.TrimSpace(m[1])
		}
	}

	// HalkBank
	if alici == "" {
		if m := reAliciHalk.FindStringSubmatch(text); len(m) > 1 {
			alici = strings.TrimSpace(m[1])
		}
	}
	if gonderen == "" {
		if m := reGonderenHalk.FindStringSubmatch(text); len(m) > 1 {
			gonderen = strings.TrimSpace(m[1])
		}
	}
	if aciklama == "" {
		if m := reAciklamaHalk.FindStringSubmatch(text); len(m) > 1 {
			aciklama = strings.TrimSpace(m[1])
		}
	}
	if tutar == "" {
		if m := reTutarHalk.FindStringSubmatch(text); len(m) > 1 {
			tutar = strings.TrimSpace(m[1])
		}
	}
	if tarih == "" {
		if m := reTarihHalk.FindStringSubmatch(text); len(m) > 1 {
			tarih = strings.TrimSpace(m[1])
		}
	}

	// Generic fallback patterns
	if alici == "" {
		if m := reAlici.FindStringSubmatch(text); len(m) > 1 {
			alici = strings.TrimSpace(m[1])
		}
	}
	if aciklama == "" {
		if m := reAciklama.FindStringSubmatch(text); len(m) > 1 {
			aciklama = strings.TrimSpace(m[1])
		}
	}
	if tutar == "" {
		if m := reTutar.FindStringSubmatch(text); len(m) > 1 {
			tutar = strings.TrimSpace(m[1])
		}
	}

	// If no specific amount field found, try a more generic approach
	if tutar == "" {
		reGenericTutar := regexp.MustCompile(`(?i)([0-9]+(?:[.,][0-9]{1,3})*(?:[.,][0-9]{2})?)\s*(?:TL|₺)`)
		matches := reGenericTutar.FindAllStringSubmatch(text, -1)
		if len(matches) > 0 {
			// Take the first found amount that looks like a transaction amount
			for _, match := range matches {
				if len(match) > 1 {
					tutar = strings.TrimSpace(match[1])
					break
				}
			}
		}
	}

	// Clean up extracted values - remove common prefixes and suffixes
	alici = cleanFieldValue(alici)
	gonderen = cleanFieldValue(gonderen)
	aciklama = cleanFieldValue(aciklama)
	tarih = cleanFieldValue(tarih)

	// Return empty string if no meaningful data was extracted
	if alici == "" && gonderen == "" && aciklama == "" && tutar == "" && tarih == "" {
		if config.EnableDebugLogging {
			p.API.LogDebug("No meaningful data extracted from PDF text",
				"textPreview", text[:min(200, len(text))],
				"author", "SkyLostTR (@Keeftraum)")
		}
		return ""
	}

	// Build the response with available information
	var result strings.Builder

	if aciklama != "" {
		result.WriteString(fmt.Sprintf("**Açıklama**: %s\n", aciklama))
	}
	if alici != "" {
		result.WriteString(fmt.Sprintf("**Alıcı**: %s\n", alici))
	}
	if gonderen != "" {
		result.WriteString(fmt.Sprintf("**Gönderen**: %s\n", gonderen))
	}
	if tutar != "" {
		result.WriteString(fmt.Sprintf("**İşlem Tutarı**: %s TL\n", tutar))
	}
	if tarih != "" {
		result.WriteString(fmt.Sprintf("**İşlem Tarihi**: %s\n", tarih))
	}

	if config.EnableDebugLogging {
		p.API.LogDebug("Successfully extracted fields",
			"fieldsCount", strings.Count(result.String(), "**"),
			"author", "SkyLostTR (@Keeftraum)")
	}

	return strings.TrimRight(result.String(), "\n")
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// cleanFieldValue removes common prefixes, suffixes and cleans up field values
// Enhanced by SkyLostTR (@Keeftraum) for better Turkish text processing
func cleanFieldValue(value string) string {
	if value == "" {
		return ""
	}

	// Remove common prefixes and suffixes
	cleanPatterns := []string{
		`^[:\-\s]+`,      // Leading colons, dashes, spaces
		`[:\-\s]+$`,      // Trailing colons, dashes, spaces
		`^(?i)(TL|₺)\s*`, // Leading currency symbols
		`\s*(?i)(TL|₺)$`, // Trailing currency symbols
		`^\d+\.\s*`,      // Leading numbers with dots (line numbers)
		`^\s*[-–—]\s*`,   // Leading dashes
		`\s*[-–—]\s*$`,   // Trailing dashes
	}

	result := value
	for _, pattern := range cleanPatterns {
		re := regexp.MustCompile(pattern)
		result = re.ReplaceAllString(result, "")
	}

	return strings.TrimSpace(result)
}

func main() {
	// PDF Dekont Parser Plugin - Developed by SkyLostTR (@Keeftraum)
	// Supporting Turkish banks with advanced PDF text extraction
	plugin.ClientMain(&Plugin{})
}
