---
name: Bank Support Request
about: Request support for a new bank's PDF format
title: '[BANK] Add support for '
labels: bank-support, enhancement
assignees: ''

---

## Bank Information
- **Bank Name**: [e.g. Türkiye İş Bankası]
- **Country**: [e.g. Turkey]
- **Bank Website**: [e.g. https://www.isbank.com.tr]

## PDF Format Details
- **Document Type**: [e.g. EFT Receipt, Wire Transfer Receipt, Account Statement]
- **Language**: [e.g. Turkish, English]
- **Typical File Size**: [e.g. 500KB - 2MB]
- **Number of Pages**: [e.g. Usually 1 page]

## Sample Fields
Please provide examples of the field names and formats as they appear in the PDF:

**Recipient/Beneficiary Field**:
```
[e.g. "ALICI ADI: John Doe" or "Beneficiary: John Smith"]
```

**Description/Reference Field**:
```
[e.g. "AÇIKLAMA: Invoice payment" or "Reference: INV-2023-001"]
```

**Amount Field**:
```
[e.g. "TUTAR: 1,500.00 TL" or "Amount: $150.00"]
```

**Other Important Fields**:
```
[List any other fields that should be extracted]
```

## Sample PDF
**Important**: Please attach a sample PDF with sensitive information removed/anonymized.
- [ ] I have removed all personal information (names, account numbers, etc.)
- [ ] I have verified the PDF still contains the field structure

## Additional Context
- Are there multiple PDF formats from this bank?
- Are there any special characters or encoding issues?
- Any other relevant information about the PDF format?

## Priority
- [ ] High (actively using this bank)
- [ ] Medium (planning to use)
- [ ] Low (nice to have)

## Checklist
- [ ] I have provided bank information
- [ ] I have described the PDF format
- [ ] I have provided field examples
- [ ] I have attached a sample PDF (anonymized)
- [ ] I have searched for existing bank support requests
