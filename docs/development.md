---
layout: default
title: Development Guide
nav_order: 5
description: "Development guide for contributing to the Mattermost PDF Dekont Parser Plugin"
---

# Development Guide
{: .no_toc }

Comprehensive guide for developers who want to contribute to or extend the Mattermost PDF Dekont Parser Plugin.
{: .fs-6 .fw-300 }

## Table of contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Development Environment Setup

### Prerequisites

Before starting development, ensure you have:

- **Go 1.19+**: [Download Go](https://golang.org/dl/)
- **Git**: For version control
- **Make**: For build automation (Windows: install via Chocolatey)
- **Mattermost Server**: For testing (development instance recommended)

### Environment Verification

```bash
# Check Go version
go version

# Check Git
git --version

# Check Make (Windows)
make --version
```

### Clone and Setup

```bash
# Clone the repository
git clone https://github.com/SkyLostTR/mattermost-dekont-plugin.git
cd mattermost-dekont-plugin

# Install dependencies
go mod download

# Verify build
make build
```

---

## Project Structure

### Core Files

```
mattermost-dekont-plugin/
├── plugin.go              # Main plugin logic
├── plugin_test.go         # Comprehensive test suite
├── plugin.json           # Plugin manifest
├── go.mod                # Go module definition
├── go.sum                # Dependency checksums
├── Makefile              # Build automation
└── README.md             # Project documentation
```

### Configuration Files

```
├── .github/
│   ├── workflows/        # CI/CD pipelines
│   ├── ISSUE_TEMPLATE/   # Issue templates
│   └── pull_request_template.md
├── .golangci.yml         # Linting configuration
├── .editorconfig         # Editor settings
└── sonar-project.properties  # Code quality
```

### Build Artifacts

```
├── dist/                 # Distribution files
├── plugin.exe           # Built plugin binary
└── *.tar.gz             # Plugin bundles
```

---

## Core Architecture

### Plugin Framework

The plugin extends Mattermost's plugin framework:

```go
type Plugin struct {
    plugin.MattermostPlugin
    // Plugin-specific fields
}

// Required plugin hooks
func (p *Plugin) OnActivate() error
func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post)
```

### PDF Processing Pipeline

1. **File Detection**: Monitor file uploads via `MessageHasBeenPosted`
2. **PDF Validation**: Verify file type and readability
3. **Text Extraction**: Use `github.com/ledongthuc/pdf` library
4. **Pattern Matching**: Apply regex patterns for field extraction
5. **Result Formatting**: Create structured output
6. **Message Update**: Modify original post with extracted data

### Key Components

#### File Handler
```go
func (p *Plugin) handleFileUpload(post *model.Post, fileInfo *model.FileInfo) error {
    // Download file content
    // Validate PDF format
    // Process PDF
    // Update post
}
```

#### PDF Processor
```go
func (p *Plugin) processPDF(content []byte) (*TransactionInfo, error) {
    // Extract text from PDF
    // Apply pattern matching
    // Clean and validate results
}
```

#### Pattern Matcher
```go
type FieldPattern struct {
    Name    string
    Regex   *regexp.Regexp
    Cleaner func(string) string
}
```

---

## Development Workflow

### Coding Standards

Follow Go best practices and project-specific standards:

#### Code Style
```go
// ✅ Good: Clear function names and error handling
func (p *Plugin) extractTransactionAmount(text string) (string, error) {
    pattern := regexp.MustCompile(`(?i)(TUTAR|İŞLEM TUTARI):\s*([0-9.,]+\s*[A-Z]+)`)
    matches := pattern.FindStringSubmatch(text)
    if len(matches) < 3 {
        return "", errors.New("amount not found")
    }
    return strings.TrimSpace(matches[2]), nil
}

// ❌ Bad: Unclear naming and poor error handling
func (p *Plugin) getAmt(s string) string {
    r, _ := regexp.Compile(`(?i)(TUTAR|İŞLEM TUTARI):\s*([0-9.,]+\s*[A-Z]+)`)
    m := r.FindStringSubmatch(s)
    return m[2]
}
```

#### Error Handling
```go
// Always log errors with context
if err != nil {
    p.API.LogError("Failed to process PDF", "error", err.Error(), "filename", fileInfo.Name)
    return err
}
```

#### Resource Management
```go
// Always clean up resources
defer func() {
    if tempFile != nil {
        os.Remove(tempFile.Name())
    }
}()
```

### Testing Strategy

#### Unit Tests
Use table-driven tests for comprehensive coverage:

```go
func TestExtractRecipient(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
        hasError bool
    }{
        {
            name:     "İş Bankası format",
            input:    "ALICI: AHMET YILMAZ",
            expected: "AHMET YILMAZ",
            hasError: false,
        },
        {
            name:     "Garanti format",
            input:    "ALAN: FATMA KAYA",
            expected: "FATMA KAYA", 
            hasError: false,
        },
        {
            name:     "No recipient found",
            input:    "Some random text",
            expected: "",
            hasError: true,
        },
    }

    p := &Plugin{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := p.extractRecipient(tt.input)
            
            if tt.hasError && err == nil {
                t.Errorf("Expected error but got none")
            }
            if !tt.hasError && err != nil {
                t.Errorf("Unexpected error: %v", err)
            }
            if result != tt.expected {
                t.Errorf("Expected %q, got %q", tt.expected, result)
            }
        })
    }
}
```

#### Benchmark Tests
```go
func BenchmarkProcessPDF(b *testing.B) {
    plugin := &Plugin{}
    pdfData := loadTestPDF("sample_receipt.pdf")
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = plugin.processPDF(pdfData)
    }
}
```

### Build Process

#### Local Development
```bash
# Build plugin
make build

# Run tests
go test -v

# Run with coverage
go test -cover

# Run benchmarks
go test -bench=.

# Create bundle
make bundle
```

#### Using Make Targets
```bash
# Clean previous builds
make clean

# Install dependencies
make deps

# Lint code
make lint

# Format code
make fmt

# Run all checks
make check
```

---

## Adding Bank Support

### Step 1: Analyze Bank Format

1. **Collect Sample PDFs**: Get various receipt types from the bank
2. **Identify Patterns**: Look for consistent field labels and formats
3. **Document Variations**: Note different ways fields might appear

### Step 2: Create Pattern Rules

```go
// Add to field extraction patterns
var recipientPatterns = []FieldPattern{
    {
        Name:  "Standard Alıcı",
        Regex: regexp.MustCompile(`(?i)ALICI[^:]*:\s*(.+)`),
        Cleaner: cleanRecipientName,
    },
    {
        Name:  "New Bank Format",
        Regex: regexp.MustCompile(`(?i)HESAP SAHİBİ[^:]*:\s*(.+)`),
        Cleaner: cleanRecipientName,
    },
}
```

### Step 3: Write Tests

```go
func TestNewBankExtraction(t *testing.T) {
    tests := []struct {
        name     string
        pdfText  string
        expected TransactionInfo
    }{
        {
            name:    "New Bank Standard Receipt",
            pdfText: loadTestPDFText("newbank_sample.txt"),
            expected: TransactionInfo{
                Recipient:   "JOHN DOE",
                Description: "Payment Description",
                Amount:     "1,000.00 TL",
            },
        },
    }
    // Test implementation...
}
```

### Step 4: Update Documentation

1. Add bank to supported list
2. Document field patterns
3. Include sample output
4. Update README.md

---

## Debugging and Testing

### Local Testing Setup

1. **Development Mattermost Instance**:
   ```bash
   # Using Docker
   docker run --name mattermost-dev -d --publish 8065:8065 mattermost/mattermost-preview
   ```

2. **Plugin Installation**:
   - Build plugin bundle
   - Upload via System Console
   - Enable plugin
   - Test with sample PDFs

### Debug Logging

Enable detailed logging:

```go
// Add debug logs throughout processing
p.API.LogDebug("Processing PDF", "filename", fileInfo.Name, "size", len(content))
p.API.LogDebug("Extracted text", "length", len(text), "preview", text[:min(100, len(text))])
p.API.LogDebug("Pattern match result", "field", "recipient", "value", recipient)
```

### Common Issues

#### PDF Text Extraction
```go
// Handle PDF extraction errors gracefully
func (p *Plugin) extractPDFText(content []byte) (string, error) {
    reader, err := pdf.NewReader(bytes.NewReader(content), int64(len(content)))
    if err != nil {
        return "", fmt.Errorf("failed to create PDF reader: %w", err)
    }
    
    var text strings.Builder
    for i := 1; i <= reader.NumPage(); i++ {
        page, err := reader.Page(i)
        if err != nil {
            p.API.LogWarn("Failed to read page", "page", i, "error", err.Error())
            continue
        }
        
        pageText, err := page.GetPlainText()
        if err != nil {
            p.API.LogWarn("Failed to extract text from page", "page", i, "error", err.Error())
            continue
        }
        
        text.WriteString(pageText)
    }
    
    return text.String(), nil
}
```

---

## Performance Optimization

### Memory Management

```go
// Process large PDFs efficiently
func (p *Plugin) processLargePDF(content []byte) error {
    // Limit PDF size
    if len(content) > maxPDFSize {
        return errors.New("PDF too large")
    }
    
    // Use streaming where possible
    reader := bytes.NewReader(content)
    
    // Clean up resources
    defer func() {
        content = nil
        runtime.GC()
    }()
    
    return nil
}
```

### Concurrent Processing

```go
// Handle multiple uploads efficiently
func (p *Plugin) processMultipleFiles(posts []*model.Post) {
    semaphore := make(chan struct{}, maxConcurrentProcessing)
    var wg sync.WaitGroup
    
    for _, post := range posts {
        wg.Add(1)
        go func(p *model.Post) {
            defer wg.Done()
            semaphore <- struct{}{}
            defer func() { <-semaphore }()
            
            p.processPost(post)
        }(post)
    }
    
    wg.Wait()
}
```

---

## CI/CD Pipeline

### GitHub Actions Workflow

The project uses automated CI/CD:

```yaml
# .github/workflows/ci.yml
name: CI
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: '1.19'
      - name: Run tests
        run: go test -v -cover ./...
      - name: Build
        run: make build
```

### Code Quality Checks

```yaml
# .github/workflows/code-quality.yml
- name: Run golangci-lint
  uses: golangci/golangci-lint-action@v3
  with:
    version: latest

- name: SonarCloud Scan
  uses: SonarSource/sonarcloud-github-action@master
```

---

## Security Guidelines

### Secure Coding Practices

1. **Input Validation**:
   ```go
   // Validate file types and sizes
   if !isValidPDF(content) {
       return errors.New("invalid PDF format")
   }
   ```

2. **No Sensitive Data Logging**:
   ```go
   // ❌ Never log sensitive information
   p.API.LogInfo("Processing transaction", "amount", amount, "recipient", recipient)
   
   // ✅ Log only metadata
   p.API.LogInfo("Processing transaction", "fieldCount", len(fields))
   ```

3. **Resource Limits**:
   ```go
   // Prevent resource exhaustion
   const (
       maxPDFSize = 10 * 1024 * 1024 // 10MB
       maxProcessingTime = 30 * time.Second
   )
   ```

### Dependency Security

- Regular dependency updates via Dependabot
- Security scanning with Trivy
- Vulnerability monitoring via GitHub Security

---

## Release Process

### Version Management

Follow semantic versioning:
- **Major**: Breaking changes
- **Minor**: New features, backward compatible
- **Patch**: Bug fixes

### Release Checklist

1. **Update Version**:
   - `plugin.json`: Update version field
   - `CHANGELOG.md`: Document changes
   - `README.md`: Update compatibility info

2. **Run Tests**:
   ```bash
   make test
   make lint
   make build
   ```

3. **Create Release**:
   - Tag version: `git tag v1.x.x`
   - Push tags: `git push --tags`
   - GitHub Actions automatically creates release

4. **Post-Release**:
   - Update documentation
   - Notify community
   - Monitor for issues

---

## Contributing Guidelines

### Pull Request Process

1. **Fork and Branch**:
   ```bash
   git checkout -b feature/new-bank-support
   ```

2. **Development**:
   - Write code following standards
   - Add comprehensive tests
   - Update documentation

3. **Testing**:
   ```bash
   make test
   make lint
   make build
   ```

4. **Submit PR**:
   - Use PR template
   - Include test results
   - Request review

### Code Review Checklist

- [ ] Tests added for new functionality
- [ ] Documentation updated
- [ ] Security considerations addressed
- [ ] Performance impact assessed
- [ ] Error handling implemented
- [ ] Turkish character encoding handled
- [ ] Bank format compatibility verified

---

## Resources

### Development Tools

- **IDE Setup**: VS Code with Go extension
- **Debugging**: Delve debugger
- **Testing**: Built-in Go testing framework
- **Profiling**: pprof for performance analysis

### Documentation

- [Mattermost Plugin API](https://developers.mattermost.com/extend/plugins/)
- [Go Best Practices](https://golang.org/doc/effective_go.html)
- [PDF Processing Library](https://github.com/ledongthuc/pdf)

### Community

- [GitHub Discussions](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)
- [Issue Tracker](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues)
- [Pull Requests](https://github.com/SkyLostTR/mattermost-dekont-plugin/pulls)

---

## Support

Need help with development?

- 💻 [Development Issues](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=development.md)
- 📖 [API Documentation](api.html)
- 🤝 [Contributing Guide](contributing.html)
- 💬 [Developer Discussions](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)
