// Package main implements a Mattermost plugin for parsing PDF bank receipts.
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

// Plugin represents the main plugin instance.
type Plugin struct {
	plugin.MattermostPlugin
}

// OnActivate is called when the plugin is activated.
func (p *Plugin) OnActivate() error {
	p.API.LogInfo("PDF Parser Plugin activated")
	return nil
}

// MessageHasBeenPosted processes newly posted messages to extract PDF content.
func (p *Plugin) MessageHasBeenPosted(_ *plugin.Context, post *model.Post) {
	if post.Type != "" || post.FileIds == nil || len(post.FileIds) == 0 {
		return
	}

	for _, fileID := range post.FileIds {
		if err := p.processFileUpload(fileID, post); err != nil {
			p.API.LogError("Failed to process file upload", "fileID", fileID, "error", err.Error())
		}
	}
}

func (p *Plugin) processFileUpload(fileID string, post *model.Post) error {
	fileInfo, err := p.API.GetFileInfo(fileID)
	if err != nil || !strings.HasSuffix(fileInfo.Name, ".pdf") {
		return nil // Not a PDF file, skip silently
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

	description := extractFields(extractedText)
	if description != "" {
		post.Message = description
		_, appErr = p.API.UpdatePost(post)
		if appErr != nil {
			return appErr
		}
	}

	return nil
}

func extractFields(text string) string {
	var alici, aciklama, tutar string

	reAlici := regexp.MustCompile(`(?i)ALICI *(?:AD SOYAD/UNVAN)? *[:\-]? *(.+?)\n`)
	reAciklama := regexp.MustCompile(`(?i)A[CÇ]IKLAMA *[:\-]? *(.+?)\n`)
	reTutar := regexp.MustCompile(`(?i)(?:IŞLEM TUTARI|TUTARI|HAVALE TUTARI|GIDEN EFT TUTARI).*?: *-?([\d.,]+) *TL`)

	if m := reAlici.FindStringSubmatch(text); len(m) > 1 {
		alici = strings.TrimSpace(m[1])
	}
	if m := reAciklama.FindStringSubmatch(text); len(m) > 1 {
		aciklama = strings.TrimSpace(m[1])
	}
	if m := reTutar.FindStringSubmatch(text); len(m) > 1 {
		tutar = strings.TrimSpace(m[1])
	}

	if alici == "" && aciklama == "" && tutar == "" {
		return ""
	}

	return fmt.Sprintf("**Açıklama**: %s\n**Alıcı**: %s\n**İşlem Tutarı**: %s TL", aciklama, alici, tutar)
}

func main() {
	plugin.ClientMain(&Plugin{})
}
