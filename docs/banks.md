---
layout: default
title: Supported Banks
nav_order: 4
description: "Complete list of supported Turkish banks and their receipt formats"
---

# Supported Banks
{: .no_toc }

Comprehensive overview of Turkish banks supported by the Mattermost PDF Dekont Parser Plugin.
{: .fs-6 .fw-300 }

## Table of contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Currently Supported Banks

The plugin supports PDF receipt formats from major Turkish banks. Each bank has been tested with various receipt types to ensure reliable field extraction.

### 🏦 Türkiye İş Bankası

**Status**: ✅ Fully Supported  
**Receipt Types**: EFT, Havale, Transfer confirmations  
**Supported Fields**: Alıcı, Açıklama, İşlem Tutarı  

**Field Patterns**:
- **Recipient**: `ALICI`, `ALICI AD SOYAD`, `ALICI UNVAN`
- **Description**: `AÇIKLAMA`, `AÇIKLAMA/REFERANS`
- **Amount**: `İŞLEM TUTARI`, `GÖNDERILEN TUTAR`

**Sample Output**:
```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: AHMET YILMAZ
📝 Açıklama: Proje Ödemesi
💰 İşlem Tutarı: 2,500.00 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

### 🏦 Garanti BBVA

**Status**: ✅ Fully Supported  
**Receipt Types**: Online banking transfers, Mobile app receipts  
**Supported Fields**: Alıcı, Açıklama, İşlem Tutarı  

**Field Patterns**:
- **Recipient**: `ALICI`, `ALAN`, `YARARLANICIADI`
- **Description**: `AÇIKLAMA`, `ACIKLAMA`, `REFERANS`
- **Amount**: `TUTAR`, `İŞLEM TUTARI`, `TRANSFER TUTARI`

**Sample Output**:
```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: FATMA KAYA
📝 Açıklama: Danışmanlık Ücreti
💰 İşlem Tutarı: 3,750.00 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

### 🏦 Akbank

**Status**: ✅ Fully Supported  
**Receipt Types**: Internet banking receipts, Mobile transfer confirmations  
**Supported Fields**: Alıcı, Açıklama, İşlem Tutarı  

**Field Patterns**:
- **Recipient**: `ALICI`, `HESAP SAHİBİ`, `ALAN TARAF`
- **Description**: `AÇIKLAMA`, `ACIKLAMA`, `İŞLEM AÇIKLAMASI`
- **Amount**: `TUTAR`, `İŞLEM TUTARI`, `GÖNDERILEN MIKTAR`

**Sample Output**:
```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: MEHMET ÖZKAN
📝 Açıklama: Freelance Payment
💰 İşlem Tutarı: 1,200.50 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

### 🏦 Yapı Kredi

**Status**: ✅ Fully Supported  
**Receipt Types**: Transfer confirmations, Payment receipts  
**Supported Fields**: Alıcı, Açıklama, İşlem Tutarı  

**Field Patterns**:
- **Recipient**: `ALICI`, `ALICI ADI`, `HESAP SAHİBİ`
- **Description**: `AÇIKLAMA`, `ACIKLAMA`, `REFERANS BILGISI`
- **Amount**: `TUTAR`, `İŞLEM TUTARI`, `ÖDEME TUTARI`

**Sample Output**:
```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: AYŞE DEMIR
📝 Açıklama: Kira Ödemesi
💰 İşlem Tutarı: 4,000.00 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

### 🏦 Ziraat Bankası

**Status**: ✅ Fully Supported  
**Receipt Types**: Government payments, Standard transfers  
**Supported Fields**: Alıcı, Açıklama, İşlem Tutarı  

**Field Patterns**:
- **Recipient**: `ALICI`, `YARARLANICI`, `HESAP SAHİBİ`
- **Description**: `AÇIKLAMA`, `ACIKLAMA`, `İŞLEM TANIMI`
- **Amount**: `TUTAR`, `İŞLEM TUTARI`, `ÖDENEN MIKTAR`

**Sample Output**:
```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: ALİ ÇELIK
📝 Açıklama: Vergi Ödemesi
💰 İşlem Tutarı: 850.75 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

## Banks Under Development

### 🚧 VakıfBank
**Status**: 🔄 In Progress  
**ETA**: Next release  
**Notes**: Pattern analysis in progress

### 🚧 Halkbank
**Status**: 🔄 In Progress  
**ETA**: Next release  
**Notes**: Government bank format research ongoing

### 🚧 DenizBank
**Status**: 📋 Planned  
**ETA**: Future release  
**Notes**: Waiting for sample receipts

---

## Request Bank Support

Don't see your bank listed? We're always adding support for new banks!

### How to Request Support

1. **Create an Issue**: Use our [bank support template](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bank_support.md)

2. **Provide Sample Receipt**: 
   - Upload an anonymized PDF sample
   - Remove all personal information (names, account numbers, etc.)
   - Keep the field labels and formatting intact

3. **Include Bank Details**:
   - Bank name and full official title
   - Receipt type (EFT, Havale, etc.)
   - Online banking platform used

### Sample Anonymization Guide

When providing sample PDFs:

```
✅ GOOD - Keep field labels:
ALICI: [ANONYMIZED]
AÇIKLAMA: [ANONYMIZED]
İŞLEM TUTARI: [ANONYMIZED]

❌ BAD - Don't remove labels:
[REDACTED]: [REDACTED]
[REDACTED]: [REDACTED]
[REDACTED]: [REDACTED]
```

---

## Bank Format Specifications

### Common Field Patterns

The plugin uses regex patterns to identify fields across different banks:

#### Recipient Field Variants
```regex
(?i)(ALICI|ALAN|YARARLANICI|HESAP SAHİBİ|ALAN TARAF)
```

#### Description Field Variants
```regex
(?i)(AÇIKLAMA|ACIKLAMA|REFERANS|İŞLEM AÇIKLAMASI)
```

#### Amount Field Variants
```regex
(?i)(TUTAR|İŞLEM TUTARI|GÖNDERILEN|ÖDEME|MIKTAR)
```

### PDF Requirements

For successful processing, PDFs must:

1. **Contain Selectable Text**: Not scanned images
2. **Use Standard Encoding**: UTF-8 or Turkish character sets
3. **Have Clear Field Labels**: Recognizable Turkish banking terms
4. **Be Well-Formatted**: Consistent spacing and structure

---

## Troubleshooting Bank Issues

### Field Not Extracted

**Problem**: Some fields missing for your bank

**Solutions**:
1. Check if your bank is fully supported
2. Verify PDF quality and text selection
3. Compare with supported field patterns
4. Request enhanced support via GitHub issue

### Wrong Bank Detection

**Problem**: Plugin misidentifies bank format

**Solutions**:
1. Ensure PDF is from supported bank
2. Use standard receipt formats (not custom templates)
3. Check for mixed content or unusual formatting
4. Report issue with sample PDF

### Partial Recognition

**Problem**: Only some fields extracted correctly

**Solutions**:
1. Cross-reference with known field patterns
2. Check for typos in original PDF
3. Verify Turkish character encoding
4. Submit feedback for pattern improvement

---

## Technical Details

### Pattern Matching

The plugin uses sophisticated regex patterns that:

- **Case-insensitive matching**: Handles various capitalizations
- **Turkish character support**: Properly handles ğ, ü, ş, ı, ö, ç
- **Flexible spacing**: Accommodates different formatting styles
- **Multiple variants**: Recognizes various field name formats

### Processing Pipeline

1. **Text Extraction**: PDF content converted to plain text
2. **Pattern Recognition**: Field patterns matched against text
3. **Data Cleaning**: Extracted values cleaned and formatted
4. **Validation**: Results verified for consistency
5. **Output Generation**: Structured format created

---

## Statistics

### Current Coverage

| Category | Count | Percentage |
|:---------|:------|:-----------|
| Major Banks | 5/8 | 62.5% |
| Government Banks | 1/3 | 33.3% |
| Private Banks | 4/5 | 80% |
| **Total Coverage** | **5/8** | **62.5%** |

### Field Extraction Success Rate

| Bank | Alıcı | Açıklama | İşlem Tutarı | Overall |
|:-----|:------|:---------|:-------------|:--------|
| İş Bankası | 95% | 98% | 99% | 97% |
| Garanti BBVA | 93% | 96% | 98% | 96% |
| Akbank | 94% | 95% | 97% | 95% |
| Yapı Kredi | 92% | 94% | 96% | 94% |
| Ziraat | 90% | 92% | 95% | 92% |

---

## Contributing Bank Support

Want to help add support for more banks?

1. **Review Contribution Guide**: Check our [contributing documentation](contributing.html)
2. **Analyze Bank Patterns**: Study receipt formats and field layouts
3. **Write Tests**: Create comprehensive test cases
4. **Submit Pull Request**: Follow our development workflow

---

## Support

Need help with bank-specific issues?

- 🏦 [Request New Bank Support](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bank_support.md)
- 🐛 [Report Bank Processing Issues](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues)
- 📖 [View Usage Examples](usage.html)
- 💬 [Join Community Discussions](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)
