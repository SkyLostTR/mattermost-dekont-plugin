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
		{
			name:     "with Turkish lira symbol",
			input:    "ALICI: Test User\nAÇIKLAMA: Payment with symbol\nTUTARI: 500.00 ₺",
			expected: "**Açıklama**: Payment with symbol\n**Alıcı**: Test User\n**İşlem Tutarı**: 500.00 TL",
		},
		{
			name:     "amount with dots as thousand separator",
			input:    "ALICI: Test User\nAÇIKLAMA: Large payment\nİŞLEM TUTARI: 1.234.567,89 TL",
			expected: "**Açıklama**: Large payment\n**Alıcı**: Test User\n**İşlem Tutarı**: 1.234.567,89 TL",
		},
		{
			name:     "amount without specific field label",
			input:    "ALICI: Test User\nAÇIKLAMA: Generic payment\nSome text 250.75 TL more text",
			expected: "**Açıklama**: Generic payment\n**Alıcı**: Test User\n**İşlem Tutarı**: 250.75 TL",
		},
		{
			name:     "transfer amount field",
			input:    "ALICI: Bank Transfer\nAÇIKLAMA: Transfer payment\nTRANSFER TUTARI: 1,000.00 TL",
			expected: "**Açıklama**: Transfer payment\n**Alıcı**: Bank Transfer\n**İşlem Tutarı**: 1,000.00 TL",
		},
		{
			name:     "amount with Turkish I character variations",
			input:    "ALICI: Test User\nAÇIKLAMA: Test payment\nİŞLEM TUTARI: 123.45 TL",
			expected: "**Açıklama**: Test payment\n**Alıcı**: Test User\n**İşlem Tutarı**: 123.45 TL",
		},
		// VakıfBank format tests
		{
			name:     "VakıfBank standard format",
			input:    "ALICI AD SOYAD/UNVAN: Mehmet Yılmaz\nGÖNDEREN AD SOYAD / UNVAN: Ahmet Kaya\nİŞLEM TUTARI: 2,500.00\nİŞLEM TARİHİ: 15.07.2025\nİŞLEM AÇIKLAMASI: Kira ödemesi",
			expected: "**Açıklama**: Kira ödemesi\n**Alıcı**: Mehmet Yılmaz\n**Gönderen**: Ahmet Kaya\n**İşlem Tutarı**: 2,500.00 TL\n**İşlem Tarihi**: 15.07.2025",
		},
		{
			name:     "VakıfBank with TL suffix",
			input:    "ALICI AD SOYAD/UNVAN: ABC Şirketi\nİŞLEM TUTARI: 1,750.50 TL\nİŞLEM AÇIKLAMASI: Fatura ödemesi",
			expected: "**Açıklama**: Fatura ödemesi\n**Alıcı**: ABC Şirketi\n**İşlem Tutarı**: 1,750.50 TL",
		},
		// YapıKredi format tests
		{
			name:     "YapıKredi standard format",
			input:    "ALICI ADI: Ayşe Demir\nGÖNDEREN ADI SOYAD: Can Özkan\nGİDEN EFT TUTARI: 3,200.75\nAÇIKLAMA: Ürün bedeli",
			expected: "**Açıklama**: Ürün bedeli\n**Alıcı**: Ayşe Demir\n**Gönderen**: Can Özkan\n**İşlem Tutarı**: 3,200.75 TL",
		},
		{
			name:     "YapıKredi with Turkish characters",
			input:    "ALICI ADI: Özgür Şahin\nAÇIKLAMA: Hizmet bedeli ödemesi\nGİDEN EFT TUTARI: 850.00 TL",
			expected: "**Açıklama**: Hizmet bedeli ödemesi\n**Alıcı**: Özgür Şahin\n**İşlem Tutarı**: 850.00 TL",
		},
		// Kuveyt Türk format tests
		{
			name:     "Kuveyt Türk standard format",
			input:    "Tutar: 1,500.25\nAçıklama: Online alışveriş\nGönderilen IBAN: TR12 3456 7890 1234 5678 90\nAlıcı: E-ticaret Mağazası\nGönderen Kişi: Fatma Yıldız",
			expected: "**Açıklama**: Online alışveriş\n**Alıcı**: E-ticaret Mağazası\n**Gönderen**: Fatma Yıldız\n**İşlem Tutarı**: 1,500.25 TL",
		},
		{
			name:     "Kuveyt Türk case insensitive",
			input:    "tutar: 750.00\naciklama: Elektrik faturası\nalici: BEDAŞ\ngönderen kişi: Hasan Çelik",
			expected: "**Açıklama**: Elektrik faturası\n**Alıcı**: BEDAŞ\n**Gönderen**: Hasan Çelik\n**İşlem Tutarı**: 750.00 TL",
		},
		// HalkBank format tests
		{
			name:     "HalkBank standard format",
			input:    "GÖNDEREN : Murat Arslan\nALICI : Teknoloji A.Ş.\nİŞLEM TUTARI (TL) : 4,250.00\nAÇIKLAMA : Yazılım lisansı\nİŞLEM TARİHİ : 20.07.2025",
			expected: "**Açıklama**: Yazılım lisansı\n**Alıcı**: Teknoloji A.Ş.\n**Gönderen**: Murat Arslan\n**İşlem Tutarı**: 4,250.00 TL\n**İşlem Tarihi**: 20.07.2025",
		},
		{
			name:     "HalkBank with special characters",
			input:    "GÖNDEREN : İrem Öztürk\nALICI : Güven Sigorta\nİŞLEM TUTARI (TL) : 1,200.50 ₺\nAÇIKLAMA : Kasko primi",
			expected: "**Açıklama**: Kasko primi\n**Alıcı**: Güven Sigorta\n**Gönderen**: İrem Öztürk\n**İşlem Tutarı**: 1,200.50 TL",
		},
		// Mixed format and edge case tests
		{
			name:     "multiple amount fields - first one wins",
			input:    "ALICI: Test User\nTUTAR: 100.00\nİŞLEM TUTARI: 200.00\nAÇIKLAMA: Test payment",
			expected: "**Açıklama**: Test payment\n**Alıcı**: Test User\n**İşlem Tutarı**: 100.00 TL",
		},
		{
			name:     "amount with comma as decimal separator",
			input:    "ALICI: Test User\nAÇIKLAMA: Payment\nTUTAR: 1.234.567,89 TL",
			expected: "**Açıklama**: Payment\n**Alıcı**: Test User\n**İşlem Tutarı**: 1.234.567,89 TL",
		},
		{
			name:     "field values with extra whitespace and colons",
			input:    "ALICI  :  Test User  \n  AÇIKLAMA  :  Payment description  \n  TUTAR  :  500.00 TL  ",
			expected: "**Açıklama**: Payment description\n**Alıcı**: Test User\n**İşlem Tutarı**: 500.00 TL",
		},
		{
			name:     "sender only (no recipient)",
			input:    "GÖNDEREN: Ali Veli\nAÇIKLAMA: Transfer\nTUTAR: 300.00",
			expected: "**Açıklama**: Transfer\n**Gönderen**: Ali Veli\n**İşlem Tutarı**: 300.00 TL",
		},
		{
			name:     "date only transaction",
			input:    "İŞLEM TARİHİ: 25.07.2025\nTUTAR: 150.00\nAÇIKLAMA: Tarihli işlem",
			expected: "**Açıklama**: Tarihli işlem\n**İşlem Tutarı**: 150.00 TL\n**İşlem Tarihi**: 25.07.2025",
		},
		{
			name:     "field value cleaning - remove line numbers",
			input:    "ALICI: 1. Test Company Ltd.\nAÇIKLAMA: 2. Service payment\nTUTAR: 1,000.00",
			expected: "**Açıklama**: Service payment\n**Alıcı**: Test Company Ltd.\n**İşlem Tutarı**: 1,000.00 TL",
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

func TestExtractFieldsNewBanks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// VakıfBank comprehensive tests
		{
			name:     "VakıfBank complete receipt",
			input:    "Sayfa 1/1\nVAKIFBANK EFT DEKONTu\nALICI AD SOYAD/UNVAN: Teknoloji Şirketi A.Ş.\nGÖNDEREN AD SOYAD / UNVAN: Mehmet Yılmaz\nİŞLEM TUTARI: 15,750.50\nİŞLEM TARİHİ: 30.07.2025 14:30:25\nİŞLEM AÇIKLAMASI: Yazılım geliştirme hizmeti bedeli\nRef No: 1234567890",
			expected: "**Açıklama**: Yazılım geliştirme hizmeti bedeli\n**Alıcı**: Teknoloji Şirketi A.Ş.\n**Gönderen**: Mehmet Yılmaz\n**İşlem Tutarı**: 15,750.50 TL\n**İşlem Tarihi**: 30.07.2025 14:30:25",
		},
		{
			name:     "VakıfBank with Turkish characters",
			input:    "ALICI AD SOYAD/UNVAN: Özgür Çelik\nGÖNDEREN AD SOYAD / UNVAN: Şükran Öztürk\nİŞLEM TUTARI: 2,350.75 TL\nİŞLEM AÇIKLAMASI: Şirket ortaklığı payı",
			expected: "**Açıklama**: Şirket ortaklığı payı\n**Alıcı**: Özgür Çelik\n**Gönderen**: Şükran Öztürk\n**İşlem Tutarı**: 2,350.75 TL",
		},
		// YapıKredi comprehensive tests
		{
			name:     "YapıKredi complete receipt",
			input:    "YAPI KREDİ BANKASI EFT DEKONTU\nALICI ADI: Güvenlik Hizmetleri Ltd. Şti.\nGÖNDEREN ADI SOYAD: Ahmet Kaya\nGİDEN EFT TUTARI: 8,900.00 TL\nAÇIKLAMA: Güvenlik hizmeti aylık bedeli\nİşlem No: YK2025073001",
			expected: "**Açıklama**: Güvenlik hizmeti aylık bedeli\n**Alıcı**: Güvenlik Hizmetleri Ltd. Şti.\n**Gönderen**: Ahmet Kaya\n**İşlem Tutarı**: 8,900.00 TL",
		},
		{
			name:     "YapıKredi minimal format",
			input:    "ALICI ADI: Fatma Demir\nGİDEN EFT TUTARI: 450.25\nAÇIKLAMA: Kişisel transfer",
			expected: "**Açıklama**: Kişisel transfer\n**Alıcı**: Fatma Demir\n**İşlem Tutarı**: 450.25 TL",
		},
		// Kuveyt Türk comprehensive tests
		{
			name:     "Kuveyt Türk complete receipt",
			input:    "KUVEYT TÜRK PARTICIPATION BANK\nTutar: 12,500.00 TL\nAçıklama: E-ticaret satış bedeli\nGönderilen IBAN: TR98 0020 5000 0000 1234 5678 90\nAlıcı: Online Mağaza Sistemi\nGönderen Kişi: Zeynep Arslan\nİşlem Referans: KT2025073001",
			expected: "**Açıklama**: E-ticaret satış bedeli\n**Alıcı**: Online Mağaza Sistemi\n**Gönderen**: Zeynep Arslan\n**İşlem Tutarı**: 12,500.00 TL",
		},
		{
			name:     "Kuveyt Türk case variations",
			input:    "TUTAR: 3,750.50\nAÇIKLAMA: Fatura ödeme\nALICI: Elektrik Dağıtım A.Ş.\nGÖNDEREN KİŞİ: Hasan Özkan",
			expected: "**Açıklama**: Fatura ödeme\n**Alıcı**: Elektrik Dağıtım A.Ş.\n**Gönderen**: Hasan Özkan\n**İşlem Tutarı**: 3,750.50 TL",
		},
		// HalkBank comprehensive tests
		{
			name:     "HalkBank complete receipt",
			input:    "HALKBANK EFT DEKONTU\nGÖNDEREN : İbrahim Yıldırım\nALICI : Medikal Cihazlar Ltd.\nİŞLEM TUTARI (TL) : 25,000.00\nAÇIKLAMA : Tıbbi cihaz alımı\nİŞLEM TARİHİ : 30.07.2025 16:45:12\nOnay Kodu: HB20250730001",
			expected: "**Açıklama**: Tıbbi cihaz alımı\n**Alıcı**: Medikal Cihazlar Ltd.\n**Gönderen**: İbrahim Yıldırım\n**İşlem Tutarı**: 25,000.00 TL\n**İşlem Tarihi**: 30.07.2025 16:45:12",
		},
		{
			name:     "HalkBank with special currency symbol",
			input:    "GÖNDEREN : Aylin Çetin\nALICI : Eğitim Kurumları A.Ş.\nİŞLEM TUTARI (TL) : 5,250.75 ₺\nAÇIKLAMA : Eğitim ücreti",
			expected: "**Açıklama**: Eğitim ücreti\n**Alıcı**: Eğitim Kurumları A.Ş.\n**Gönderen**: Aylin Çetin\n**İşlem Tutarı**: 5,250.75 TL",
		},
		// Mixed bank detection tests
		{
			name:     "multiple bank patterns - VakıfBank priority",
			input:    "ALICI AD SOYAD/UNVAN: VakıfBank Alıcı\nALICI ADI: YapıKredi Alıcı\nİŞLEM TUTARI: 1,000.00\nGİDEN EFT TUTARI: 2,000.00\nİŞLEM AÇIKLAMASI: VakıfBank işlemi",
			expected: "**Açıklama**: VakıfBank işlemi\n**Alıcı**: VakıfBank Alıcı\n**İşlem Tutarı**: 1,000.00 TL",
		},
		// Fallback to generic patterns
		{
			name:     "generic pattern fallback",
			input:    "ALICI: Generic Bank Alıcı\nAÇIKLAMA: Generic işlem\nİŞLEM TUTARI: 500.00 TL",
			expected: "**Açıklama**: Generic işlem\n**Alıcı**: Generic Bank Alıcı\n**İşlem Tutarı**: 500.00 TL",
		},
		// Error resilient parsing
		{
			name:     "malformed field separators",
			input:    "ALICI ADI-- Broken Format User\nAÇIKLAMA === Weird separators\nTUTAR>>> 750.00 TL",
			expected: "**Açıklama**: Weird separators\n**Alıcı**: Broken Format User\n**İşlem Tutarı**: 750.00 TL",
		},
		{
			name:     "amount without TL suffix but with Turkish lira symbol",
			input:    "ALICI: Currency Test\nAÇIKLAMA: Symbol test\nTUTAR: 1,500.50 ₺",
			expected: "**Açıklama**: Symbol test\n**Alıcı**: Currency Test\n**İşlem Tutarı**: 1,500.50 TL",
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

func TestCleanFieldValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "remove leading colons and spaces",
			input:    ":  Test Value",
			expected: "Test Value",
		},
		{
			name:     "remove trailing colons and dashes",
			input:    "Test Value  :-",
			expected: "Test Value",
		},
		{
			name:     "remove line numbers",
			input:    "1. Company Name Ltd.",
			expected: "Company Name Ltd.",
		},
		{
			name:     "remove currency symbols",
			input:    "TL 1000.00 ₺",
			expected: "1000.00",
		},
		{
			name:     "remove dashes",
			input:    "- Test Company -",
			expected: "Test Company",
		},
		{
			name:     "complex cleaning",
			input:    ": 2. - Test Value TL :",
			expected: "Test Value",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "only special characters",
			input:    ":- TL ₺",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanFieldValue(tt.input)
			if result != tt.expected {
				t.Errorf("cleanFieldValue() = %q, want %q", result, tt.expected)
			}
		})
	}
}
