package main

import (
	"strings"
	"testing"
)

func TestExtractFields(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid Turkish bank receipt",
			input:    "ALICI AD SOYAD/UNVAN: John Doe\nAÇIKLAMA: Invoice payment\nIŞLEM TUTARI: 1,500.00 TL",
			expected: "**Açıklama**: Invoice payment\n**Alıcı**: John Doe\n**İşlem Tutarı**: 1,500.00 TL",
		},
		{
			name:     "receipt with alternative amount field",
			input:    "ALICI: Jane Smith\nAÇIKLAMA: Rent payment\nHAVALE TUTARI: 2,000.50 TL",
			expected: "**Açıklama**: Rent payment\n**Alıcı**: Jane Smith\n**İşlem Tutarı**: 2,000.50 TL",
		},
		{
			name:     "receipt with EFT amount field",
			input:    "ALICI: ABC Company\nAÇIKLAMA: Service fee\nGIDEN EFT TUTARI: 750.25 TL",
			expected: "**Açıklama**: Service fee\n**Alıcı**: ABC Company\n**İşlem Tutarı**: 750.25 TL",
		},
		{
			name:     "receipt with Turkish characters",
			input:    "ALICI: Müşteri Adı\nAÇIKLAMA: Ödeme açıklaması\nTUTARI: 500.00 TL",
			expected: "**Açıklama**: Ödeme açıklaması\n**Alıcı**: Müşteri Adı\n**İşlem Tutarı**: 500.00 TL",
		},
		{
			name:     "receipt with missing recipient",
			input:    "AÇIKLAMA: Payment without recipient\nTUTARI: 100.00 TL",
			expected: "**Açıklama**: Payment without recipient\n**Alıcı**: \n**İşlem Tutarı**: 100.00 TL",
		},
		{
			name:     "receipt with missing description",
			input:    "ALICI: John Doe\nTUTARI: 200.00 TL",
			expected: "**Açıklama**: \n**Alıcı**: John Doe\n**İşlem Tutarı**: 200.00 TL",
		},
		{
			name:     "receipt with missing amount",
			input:    "ALICI: John Doe\nAÇIKLAMA: Test payment\n",
			expected: "**Açıklama**: Test payment\n**Alıcı**: John Doe\n**İşlem Tutarı**:  TL",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "no matching fields",
			input:    "Some random text without proper fields",
			expected: "",
		},
		{
			name:     "case insensitive matching",
			input:    "alici: lowercase test\naciklama: Lower case desc\ntutari: 50.00 TL",
			expected: "**Açıklama**: Lower case desc\n**Alıcı**: lowercase test\n**İşlem Tutarı**: 50.00 TL",
		},
		{
			name:     "with special characters in amount",
			input:    "ALICI: Test User\nAÇIKLAMA: Special payment\nTUTARI: -1,234.56 TL",
			expected: "**Açıklama**: Special payment\n**Alıcı**: Test User\n**İşlem Tutarı**: 1,234.56 TL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractFields(tt.input)
			if result != tt.expected {
				t.Errorf("extractFields() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestExtractFieldsEdgeCases(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		hasContent bool
	}{
		{
			name:       "multiline recipient name",
			input:      "ALICI: Very Long Company Name\nThat Spans Multiple Lines\nAÇIKLAMA: Payment\nTUTARI: 100.00 TL",
			hasContent: true,
		},
		{
			name:       "extra whitespace",
			input:      "  ALICI  :   John Doe   \n  AÇIKLAMA  :   Payment   \n  TUTARI  :   100.00 TL  ",
			hasContent: true,
		},
		{
			name:       "different line endings",
			input:      "ALICI: John Doe\r\nAÇIKLAMA: Payment\r\nTUTARI: 100.00 TL\r\n",
			hasContent: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractFields(tt.input)
			hasContent := result != ""
			if hasContent != tt.hasContent {
				t.Errorf("extractFields() returned content = %v, want %v", hasContent, tt.hasContent)
			}
			if hasContent {
				// Check that result contains expected markers
				if !strings.Contains(result, "**Açıklama**:") ||
					!strings.Contains(result, "**Alıcı**:") ||
					!strings.Contains(result, "**İşlem Tutarı**:") {
					t.Errorf("extractFields() result missing expected format markers: %q", result)
				}
			}
		})
	}
}

func BenchmarkExtractFields(b *testing.B) {
	input := "ALICI AD SOYAD/UNVAN: John Doe\nAÇIKLAMA: Invoice payment\nIŞLEM TUTARI: 1,500.00 TL"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		extractFields(input)
	}
}

func BenchmarkExtractFieldsLargeInput(b *testing.B) {
	// Simulate a large PDF with lots of text
	input := strings.Repeat("Random text line\n", 1000) +
		"ALICI AD SOYAD/UNVAN: John Doe\nAÇIKLAMA: Invoice payment\nIŞLEM TUTARI: 1,500.00 TL" +
		strings.Repeat("\nMore random text", 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		extractFields(input)
	}
}
