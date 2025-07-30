---
layout: default
title: Usage Guide
nav_order: 3
description: "How to use the Mattermost PDF Dekont Parser Plugin effectively"
---

# Usage Guide
{: .no_toc }

Learn how to effectively use the Mattermost PDF Dekont Parser Plugin to extract transaction details from bank receipts.
{: .fs-6 .fw-300 }

## Table of contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Basic Usage

The plugin works automatically once installed and enabled. Here's how to use it:

### Step 1: Upload PDF Receipt

1. **Open any Mattermost channel** where you want to share the transaction info
2. **Upload a PDF file** using one of these methods:
   - Drag and drop the PDF file into the message area
   - Click the attachment icon (📎) and select your PDF
   - Copy and paste the PDF file directly

### Step 2: Automatic Processing

The plugin automatically:
1. **Detects** the uploaded PDF file
2. **Extracts** text content from the PDF
3. **Parses** transaction details using pattern recognition
4. **Updates** your message with structured information

### Step 3: View Results

Your original message gets updated with extracted information:

```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: AHMET YILMAZ
📝 Açıklama: Freelance Ödeme
💰 İşlem Tutarı: 2,500.00 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

## Supported Fields

The plugin extracts the following transaction fields:

### 👤 Alıcı (Recipient/Beneficiary)
- **Turkish variants**: ALICI, ALICI AD SOYAD/UNVAN, ALICI IBAN
- **Description**: Name of the person or organization receiving the payment
- **Example**: "JOHN DOE", "ABC YAZILIM LTD ŞTİ"

### 📝 Açıklama (Description/Reference)
- **Turkish variants**: AÇIKLAMA, ACIKLAMA, AÇIKLAMA/REFERANS
- **Description**: Transaction description or reference information
- **Example**: "Freelance Payment", "Invoice #12345"

### 💰 İşlem Tutarı (Transaction Amount)
- **Turkish variants**: İŞLEM TUTARI, TUTAR, GÖNDERILEN TUTAR
- **Description**: The monetary amount of the transaction
- **Example**: "1,500.00 TL", "€250.00", "$100.00"

---

## Bank-Specific Examples

### Türkiye İş Bankası
```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: MEHMET KAYA
📝 Açıklama: Kira Ödemesi
💰 İşlem Tutarı: 3,000.00 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

### Garanti BBVA
```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: AYŞE DEMIR
📝 Açıklama: Danışmanlık Ücreti
💰 İşlem Tutarı: 1,750.50 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

### Akbank
```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: ALİ YILDIRIM
📝 Açıklama: Proje Ödemesi
💰 İşlem Tutarı: 5,000.00 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

## Best Practices

### For Better Recognition

1. **Use High-Quality PDFs**:
   - Download PDFs directly from your bank's online portal
   - Avoid scanned documents when possible
   - Ensure text is selectable in the PDF

2. **Check Bank Compatibility**:
   - Verify your bank is in the [supported banks list](banks.html)
   - Use standard transfer receipt formats
   - Avoid custom or modified receipt templates

3. **File Naming**:
   - Use descriptive filenames: `isbank_transfer_20241201.pdf`
   - Include date and bank name for easier tracking
   - Avoid special characters in filenames

### Channel Organization

1. **Dedicated Channels**:
   - Create specific channels for financial documents
   - Use clear naming: `#finance-receipts`, `#bank-transfers`
   - Set appropriate channel permissions

2. **Message Context**:
   - Add context before uploading: "Payment for Project X"
   - Use thread replies for discussions about transactions
   - Tag relevant team members when needed

---

## Advanced Features

### Batch Processing

You can upload multiple PDF receipts in a single message:

1. Select multiple PDF files when uploading
2. Each PDF will be processed individually
3. Results appear as separate formatted blocks

### Error Handling

When processing fails, you'll see helpful error messages:

```
❌ PDF İşleme Hatası:
Bu PDF dosyası desteklenen bir banka formatında değil.
Lütfen farklı bir dosya deneyin veya banka desteği talep edin.
```

### Debug Information

Enable debug logging in System Console for detailed processing info:
- Extraction progress
- Pattern matching results
- Error diagnostics

---

## Troubleshooting Usage Issues

### PDF Not Being Processed

**Problem**: PDF uploaded but no extraction occurs

**Solutions**:
1. Check file is actually a PDF (not renamed image)
2. Verify PDF contains selectable text
3. Ensure file size is reasonable (<10MB)
4. Try re-uploading the file

### Incomplete Field Extraction

**Problem**: Some fields missing from output

**Solutions**:
1. Check if your bank format is fully supported
2. Verify PDF quality and text clarity
3. [Request enhanced bank support](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bank_support.md)

### Wrong Field Values

**Problem**: Extracted values don't match PDF content

**Solutions**:
1. Check for typos or formatting issues in original PDF
2. Verify Turkish character encoding
3. Report the issue with sample PDF (anonymized)

---

## Privacy and Security

### Data Handling
- **No Data Storage**: Extracted information is not stored permanently
- **Memory Processing**: PDFs processed in memory only
- **Temporary Files**: Automatically cleaned up after processing
- **No External Calls**: All processing happens locally

### Sensitive Information
- **Anonymize PDFs**: Remove personal details before uploading
- **Use Private Channels**: Share financial documents in appropriate channels
- **Regular Cleanup**: Periodically clean up financial message history

---

## Tips and Tricks

### Workflow Integration

1. **Accounting Integration**:
   - Copy extracted data for bookkeeping software
   - Use consistent formatting for easier tracking
   - Create templates for common transaction types

2. **Team Collaboration**:
   - Use extracted data in project discussions
   - Reference transaction details in reports
   - Share payment confirmations with stakeholders

3. **Record Keeping**:
   - Keep original PDFs in dedicated file storage
   - Use extracted data for quick reference
   - Maintain transaction logs based on extracted info

---

## Next Steps

- [Explore supported banks](banks.html)
- [Learn about development](development.html)
- [Contribute to the project](contributing.html)

---

## Support

Having trouble using the plugin?

- 🐛 [Report Usage Issues](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues)
- 💡 [Request Features](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=feature_request.md)
- 📖 [Check FAQ](https://github.com/SkyLostTR/mattermost-dekont-plugin/wiki/FAQ)
- 💬 [Community Discussions](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)
