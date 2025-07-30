---
layout: default
title: API Reference
nav_order: 6
description: "Complete API reference for the Mattermost PDF Dekont Parser Plugin"
---

# API Reference
{: .no_toc }

Complete technical reference for the Mattermost PDF Dekont Parser Plugin API and internal functions.
{: .fs-6 .fw-300 }

## Table of contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Plugin Interface

### Core Plugin Structure

```go
type Plugin struct {
    plugin.MattermostPlugin
    
    // Internal state
    configuration atomic.Value
}
```

### Required Plugin Hooks

#### OnActivate
```go
func (p *Plugin) OnActivate() error
```

**Description**: Called when the plugin is activated. Initializes the plugin state and validates configuration.

**Returns**: `error` - nil if activation successful, error otherwise

**Example**:
```go
func (p *Plugin) OnActivate() error {
    p.API.LogInfo("PDF Dekont Parser Plugin activated")
    
    // Initialize configuration
    if err := p.loadConfiguration(); err != nil {
        return fmt.Errorf("failed to load configuration: %w", err)
    }
    
    return nil
}
```

#### MessageHasBeenPosted
```go
func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) *model.Post
```

**Description**: Called when a message is posted. Checks for PDF file attachments and processes them.

**Parameters**:
- `c *plugin.Context` - Plugin execution context
- `post *model.Post` - The posted message

**Returns**: `*model.Post` - Modified post with extracted information (if applicable)

**Example**:
```go
func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) *model.Post {
    if len(post.FileIds) == 0 {
        return nil
    }
    
    for _, fileId := range post.FileIds {
        if err := p.processFileIfPDF(post, fileId); err != nil {
            p.API.LogError("Failed to process file", "error", err.Error())
        }
    }
    
    return nil
}
```

---

## Core Functions

### File Processing

#### processFileIfPDF
```go
func (p *Plugin) processFileIfPDF(post *model.Post, fileId string) error
```

**Description**: Checks if a file is a PDF and processes it for transaction data extraction.

**Parameters**:
- `post *model.Post` - The message post containing the file
- `fileId string` - Unique identifier of the uploaded file

**Returns**: `error` - nil if processing successful, error otherwise

**Example Usage**:
```go
err := p.processFileIfPDF(post, fileId)
if err != nil {
    p.API.LogError("PDF processing failed", "error", err.Error())
}
```

#### downloadFile
```go
func (p *Plugin) downloadFile(fileId string) ([]byte, *model.FileInfo, error)
```

**Description**: Downloads file content from Mattermost file storage.

**Parameters**:
- `fileId string` - File identifier

**Returns**:
- `[]byte` - File content as byte array
- `*model.FileInfo` - File metadata
- `error` - Download error if any

### PDF Processing

#### processPDF
```go
func (p *Plugin) processPDF(content []byte) (*TransactionInfo, error)
```

**Description**: Main PDF processing function that extracts text and parses transaction fields.

**Parameters**:
- `content []byte` - PDF file content

**Returns**:
- `*TransactionInfo` - Extracted transaction information
- `error` - Processing error if any

**Example**:
```go
pdfData, err := ioutil.ReadFile("sample.pdf")
if err != nil {
    return err
}

info, err := p.processPDF(pdfData)
if err != nil {
    return fmt.Errorf("PDF processing failed: %w", err)
}

fmt.Printf("Recipient: %s\n", info.Recipient)
```

#### extractTextFromPDF
```go
func (p *Plugin) extractTextFromPDF(content []byte) (string, error)
```

**Description**: Extracts plain text content from PDF using the pdf library.

**Parameters**:
- `content []byte` - PDF file content

**Returns**:
- `string` - Extracted text content
- `error` - Extraction error if any

---

## Data Structures

### TransactionInfo
```go
type TransactionInfo struct {
    Recipient   string `json:"recipient"`   // Alıcı
    Description string `json:"description"` // Açıklama  
    Amount      string `json:"amount"`      // İşlem Tutarı
}
```

**Description**: Structure containing extracted transaction information.

**Fields**:
- `Recipient` - Name of payment recipient (Alıcı)
- `Description` - Transaction description or reference (Açıklama)
- `Amount` - Transaction amount with currency (İşlem Tutarı)

**Example**:
```go
info := &TransactionInfo{
    Recipient:   "AHMET YILMAZ",
    Description: "Freelance Payment", 
    Amount:      "2,500.00 TL",
}
```

### FieldPattern
```go
type FieldPattern struct {
    Name    string
    Regex   *regexp.Regexp
    Cleaner func(string) string
}
```

**Description**: Pattern definition for field extraction.

**Fields**:
- `Name` - Human-readable pattern name
- `Regex` - Compiled regular expression for matching
- `Cleaner` - Function to clean extracted values

---

## Field Extraction Functions

### extractRecipient
```go
func (p *Plugin) extractRecipient(text string) (string, error)
```

**Description**: Extracts recipient information from PDF text using predefined patterns.

**Parameters**:
- `text string` - PDF text content

**Returns**:
- `string` - Recipient name
- `error` - Extraction error if not found

**Supported Patterns**:
- `ALICI:` - Standard recipient field
- `ALAN:` - Alternative recipient format
- `YARARLANICI:` - Beneficiary format
- `HESAP SAHİBİ:` - Account holder format

**Example**:
```go
text := "ALICI: MEHMET KAYA\nTUTAR: 1000 TL"
recipient, err := p.extractRecipient(text)
if err == nil {
    fmt.Println("Recipient:", recipient) // Output: MEHMET KAYA
}
```

### extractDescription
```go
func (p *Plugin) extractDescription(text string) (string, error)
```

**Description**: Extracts transaction description from PDF text.

**Parameters**:
- `text string` - PDF text content

**Returns**:
- `string` - Transaction description
- `error` - Extraction error if not found

**Supported Patterns**:
- `AÇIKLAMA:` - Standard description field
- `ACIKLAMA:` - Alternative spelling
- `REFERANS:` - Reference information
- `İŞLEM AÇIKLAMASI:` - Transaction description

### extractAmount
```go
func (p *Plugin) extractAmount(text string) (string, error)
```

**Description**: Extracts transaction amount from PDF text.

**Parameters**:
- `text string` - PDF text content

**Returns**:
- `string` - Transaction amount with currency
- `error` - Extraction error if not found

**Supported Patterns**:
- `İŞLEM TUTARI:` - Transaction amount
- `TUTAR:` - Amount
- `GÖNDERILEN TUTAR:` - Sent amount
- `ÖDEME TUTARI:` - Payment amount

---

## Utility Functions

### cleanText
```go
func (p *Plugin) cleanText(text string) string
```

**Description**: Cleans extracted text by removing extra whitespace and formatting.

**Parameters**:
- `text string` - Raw extracted text

**Returns**:
- `string` - Cleaned text

**Cleaning Operations**:
- Trims leading/trailing whitespace
- Normalizes internal whitespace
- Removes special characters where appropriate
- Handles Turkish character encoding

### formatTransactionInfo
```go
func (p *Plugin) formatTransactionInfo(info *TransactionInfo) string
```

**Description**: Formats transaction information into a structured display format.

**Parameters**:
- `info *TransactionInfo` - Transaction data

**Returns**:
- `string` - Formatted message text

**Example Output**:
```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: AHMET YILMAZ
📝 Açıklama: Freelance Payment
💰 İşlem Tutarı: 2,500.00 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

### isPDF
```go
func (p *Plugin) isPDF(fileInfo *model.FileInfo) bool
```

**Description**: Checks if a file is a PDF based on MIME type and extension.

**Parameters**:
- `fileInfo *model.FileInfo` - File metadata

**Returns**:
- `bool` - true if file is PDF, false otherwise

---

## Error Handling

### Error Types

#### PDFProcessingError
```go
type PDFProcessingError struct {
    Operation string
    Filename  string
    Err       error
}

func (e *PDFProcessingError) Error() string {
    return fmt.Sprintf("PDF processing failed during %s for file %s: %v", 
        e.Operation, e.Filename, e.Err)
}
```

#### FieldExtractionError
```go
type FieldExtractionError struct {
    Field string
    Err   error
}

func (e *FieldExtractionError) Error() string {
    return fmt.Sprintf("failed to extract %s: %v", e.Field, e.Err)
}
```

### Error Handling Patterns

```go
// Graceful error handling with logging
func (p *Plugin) processWithErrorHandling(content []byte, filename string) error {
    defer func() {
        if r := recover(); r != nil {
            p.API.LogError("Panic during PDF processing", 
                "filename", filename, "panic", r)
        }
    }()
    
    info, err := p.processPDF(content)
    if err != nil {
        return &PDFProcessingError{
            Operation: "text extraction",
            Filename:  filename,
            Err:       err,
        }
    }
    
    if info.Recipient == "" {
        return &FieldExtractionError{
            Field: "recipient",
            Err:   errors.New("no recipient pattern matched"),
        }
    }
    
    return nil
}
```

---

## Configuration

### Plugin Configuration
```go
type Configuration struct {
    DebugLogging bool `json:"debug_logging"`
    MaxFileSize  int  `json:"max_file_size"`
}
```

**Description**: Plugin configuration structure.

**Fields**:
- `DebugLogging` - Enable detailed debug logging
- `MaxFileSize` - Maximum PDF file size in bytes

### Configuration Functions

#### loadConfiguration
```go
func (p *Plugin) loadConfiguration() error
```

**Description**: Loads plugin configuration from Mattermost settings.

#### getConfiguration
```go
func (p *Plugin) getConfiguration() *Configuration
```

**Description**: Returns current plugin configuration.

---

## Constants

### File Processing
```go
const (
    MaxPDFSize         = 10 * 1024 * 1024 // 10MB
    MaxProcessingTime  = 30 * time.Second
    MaxPagesToProcess  = 10
)
```

### Field Patterns
```go
const (
    RecipientPattern    = `(?i)(ALICI|ALAN|YARARLANICI|HESAP SAHİBİ)[^:]*:\s*(.+)`
    DescriptionPattern  = `(?i)(AÇIKLAMA|ACIKLAMA|REFERANS)[^:]*:\s*(.+)`
    AmountPattern      = `(?i)(TUTAR|İŞLEM TUTARI|GÖNDERILEN)[^:]*:\s*([0-9.,]+\s*[A-Z]+)`
)
```

---

## Performance Metrics

### Benchmarking Functions

#### BenchmarkPDFProcessing
```go
func BenchmarkPDFProcessing(b *testing.B) {
    plugin := &Plugin{}
    testData := loadTestPDF()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = plugin.processPDF(testData)
    }
}
```

**Expected Performance**:
- Small PDFs (< 1MB): < 500ms
- Medium PDFs (1-5MB): < 2s
- Large PDFs (5-10MB): < 5s

### Memory Usage

#### Memory-Efficient Processing
```go
func (p *Plugin) processLargePDF(content []byte) error {
    // Process in chunks to manage memory
    if len(content) > MaxPDFSize {
        return errors.New("PDF too large")
    }
    
    // Use streaming for large files
    reader := bytes.NewReader(content)
    
    // Clean up explicitly
    defer func() {
        content = nil
        runtime.GC()
    }()
    
    return nil
}
```

---

## Testing API

### Test Helpers

#### LoadTestPDF
```go
func LoadTestPDF(filename string) []byte
```

**Description**: Loads test PDF data for unit tests.

#### CreateMockPlugin
```go
func CreateMockPlugin() *Plugin
```

**Description**: Creates a mock plugin instance for testing.

#### AssertTransactionInfo
```go
func AssertTransactionInfo(t *testing.T, expected, actual *TransactionInfo)
```

**Description**: Asserts transaction info equality in tests.

### Test Data
```go
var TestPDFs = map[string]string{
    "isbank_sample":   "testdata/isbank_receipt.pdf",
    "garanti_sample":  "testdata/garanti_receipt.pdf", 
    "akbank_sample":   "testdata/akbank_receipt.pdf",
}
```

---

## Integration Examples

### Custom Field Extraction
```go
// Add custom field extraction
func (p *Plugin) extractCustomField(text, pattern string) (string, error) {
    regex := regexp.MustCompile(pattern)
    matches := regex.FindStringSubmatch(text)
    if len(matches) < 2 {
        return "", errors.New("field not found")
    }
    return p.cleanText(matches[1]), nil
}
```

### Webhook Integration
```go
// Send extracted data to external webhook
func (p *Plugin) sendToWebhook(info *TransactionInfo) error {
    payload, _ := json.Marshal(info)
    
    resp, err := http.Post(webhookURL, "application/json", 
        bytes.NewBuffer(payload))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    return nil
}
```

---

## Support

For API-related questions:

- 📖 [Development Guide](development.html)
- 🐛 [Report API Issues](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues)
- 💬 [API Discussions](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)
- 📧 [Contact Developer](mailto:keeftraum@protonmail.com)
