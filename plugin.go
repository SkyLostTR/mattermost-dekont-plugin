package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) OnActivate() error {
	p.API.LogInfo("PDF Parser Plugin activated")
	return nil
}

func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	if post.Type != "" || post.FileIds == nil || len(post.FileIds) == 0 {
		return
	}

	for _, fileID := range post.FileIds {
		fileInfo, err := p.API.GetFileInfo(fileID)
		if err != nil || !strings.HasSuffix(fileInfo.Name, ".pdf") {
			continue
		}

		data, appErr := p.API.GetFile(fileID)
		if appErr != nil {
			p.API.LogError("Failed to get file", "error", appErr.Error())
			continue
		}

		tempFile, fileErr := ioutil.TempFile("", "*.pdf")
		if fileErr != nil {
			p.API.LogError("Failed to create temp file", "error", fileErr.Error())
			continue
		}
		defer os.Remove(tempFile.Name())

		tempFile.Write(data)
		tempFile.Close()

		file, r, pdfErr := pdf.Open(tempFile.Name())
		if pdfErr != nil {
			p.API.LogError("Failed to open PDF", "error", pdfErr.Error())
			continue
		}
		defer file.Close()

		var extractedText string
		if r.NumPage() > 0 {
			page := r.Page(1)
			if page.V.IsNull() {
				continue
			}
			extractedText, pdfErr = page.GetPlainText(nil)
			if pdfErr != nil {
				p.API.LogError("Failed to extract text from PDF", "error", pdfErr.Error())
				continue
			}
		}

		description := extractFields(extractedText)
		if description != "" {
			post.Message = description
			_, appErr = p.API.UpdatePost(post)
			if appErr != nil {
				p.API.LogError("Failed to update post", "error", appErr.Error())
			}
		}
	}
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
