---
layout: default
title: Contributing
nav_order: 7
description: "How to contribute to the Mattermost PDF Dekont Parser Plugin project"
---

# Contributing to Mattermost PDF Dekont Parser Plugin
{: .no_toc }

We love your input! We want to make contributing to this project as easy and transparent as possible.
{: .fs-6 .fw-300 }

## Table of contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Ways to Contribute

### 🐛 Report Bugs
Found a bug? Help us improve by reporting it!

**Before submitting**:
- Check existing [issues](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues)
- Make sure you're using the latest version
- Gather relevant information (logs, screenshots, PDF samples)

**Use our bug report template**:
[🐛 Report Bug](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bug_report.md){: .btn .btn-outline }

### 💡 Request Features
Have an idea for improvement? We'd love to hear it!

**Feature request guidelines**:
- Describe the problem you're trying to solve
- Explain your proposed solution
- Consider alternative approaches
- Think about potential implementation challenges

[💡 Request Feature](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=feature_request.md){: .btn .btn-outline }

### 🏦 Request Bank Support
Need support for a new bank? Help us add it!

**What we need**:
- Bank name and official title
- Sample PDF receipts (anonymized)
- Receipt type information (EFT, Havale, etc.)
- Field format documentation

[🏦 Request Bank Support](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bank_support.md){: .btn .btn-outline }

### 📝 Improve Documentation
Help make our documentation better!

**Documentation needs**:
- Fix typos and grammar
- Add examples and use cases
- Improve clarity and organization
- Translate to other languages

### 💻 Contribute Code
Ready to dive into the code? Great!

**Areas needing help**:
- Bank format support
- Error handling improvements
- Performance optimizations
- Testing coverage
- UI/UX enhancements

---

## Development Workflow

### 1. Fork and Clone

```bash
# Fork the repository on GitHub
# Then clone your fork
git clone https://github.com/YOUR_USERNAME/mattermost-dekont-plugin.git
cd mattermost-dekont-plugin

# Add upstream remote
git remote add upstream https://github.com/SkyLostTR/mattermost-dekont-plugin.git
```

### 2. Set Up Development Environment

```bash
# Install dependencies
go mod download

# Verify setup
go version  # Should be 1.19+
make build  # Should complete successfully
```

### 3. Create Feature Branch

```bash
# Sync with upstream
git fetch upstream
git checkout main
git merge upstream/main

# Create feature branch
git checkout -b feature/your-feature-name
```

### 4. Make Changes

Follow our [coding standards](#coding-standards) while making changes:

- Write clear, self-documenting code
- Add tests for new functionality
- Update documentation as needed
- Follow Go conventions and best practices

### 5. Test Your Changes

```bash
# Run tests
go test -v ./...

# Run with coverage
go test -cover ./...

# Run linting
make lint

# Test build
make build
```

### 6. Commit Changes

Follow [Conventional Commits](https://www.conventionalcommits.org/) format:

```bash
# Examples of good commit messages
git commit -m "feat: add support for VakıfBank receipts"
git commit -m "fix: handle PDFs with missing recipient field"
git commit -m "docs: update installation guide for v1.1.0"
git commit -m "test: add benchmarks for large PDF processing"
```

**Commit Types**:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code formatting (no logic changes)
- `refactor`: Code restructuring (no feature changes)
- `test`: Adding or updating tests
- `perf`: Performance improvements
- `ci`: CI/CD changes

### 7. Push and Create Pull Request

```bash
# Push your branch
git push origin feature/your-feature-name

# Create PR via GitHub web interface
```

---

## Coding Standards

### Go Code Style

#### Naming Conventions
```go
// ✅ Good: Clear, descriptive names
func extractRecipientName(text string) (string, error) {}
var recipientPatterns []FieldPattern
const MaxPDFFileSize = 10 * 1024 * 1024

// ❌ Bad: Unclear, abbreviated names
func getR(s string) string {}
var rp []FP
const MAX_SIZE = 10485760
```

#### Error Handling
```go
// ✅ Good: Comprehensive error handling
func (p *Plugin) processPDF(content []byte) (*TransactionInfo, error) {
    if len(content) == 0 {
        return nil, errors.New("empty PDF content")
    }
    
    text, err := p.extractTextFromPDF(content)
    if err != nil {
        return nil, fmt.Errorf("text extraction failed: %w", err)
    }
    
    // Process text...
    return info, nil
}

// ❌ Bad: Ignoring errors
func (p *Plugin) processPDF(content []byte) *TransactionInfo {
    text, _ := p.extractTextFromPDF(content)
    // Process without error checking...
    return info
}
```

#### Function Design
```go
// ✅ Good: Single responsibility, clear interface
func (p *Plugin) extractRecipient(text string) (string, error) {
    pattern := regexp.MustCompile(`(?i)ALICI[^:]*:\s*(.+)`)
    matches := pattern.FindStringSubmatch(text)
    if len(matches) < 2 {
        return "", errors.New("recipient not found")
    }
    return strings.TrimSpace(matches[1]), nil
}

// ❌ Bad: Multiple responsibilities, unclear purpose
func (p *Plugin) processStuff(text string) (string, string, string, error) {
    // Extract recipient, description, amount all in one function
    // Complex logic mixed together
}
```

### Documentation Standards

#### Function Documentation
```go
// extractRecipient extracts the recipient name from PDF text using 
// predefined patterns for Turkish bank receipts.
//
// The function searches for common recipient field labels like "ALICI",
// "ALAN", and "YARARLANICI" and returns the associated value.
//
// Parameters:
//   - text: The extracted PDF text content
//
// Returns:
//   - string: The recipient name if found
//   - error: Error if no recipient pattern matches
//
// Example:
//   recipient, err := p.extractRecipient("ALICI: AHMET YILMAZ")
//   // recipient = "AHMET YILMAZ", err = nil
func (p *Plugin) extractRecipient(text string) (string, error) {
    // Implementation...
}
```

#### Package Documentation
```go
// Package dekont provides PDF parsing functionality for Turkish bank receipts.
//
// This package implements a Mattermost plugin that automatically detects
// and processes PDF bank receipt uploads, extracting key transaction
// information such as recipient, description, and amount.
//
// Supported Banks:
//   - Türkiye İş Bankası
//   - Garanti BBVA
//   - Akbank
//   - Yapı Kredi
//   - Ziraat Bankası
//
// Example usage:
//   plugin := &Plugin{}
//   info, err := plugin.processPDF(pdfContent)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Printf("Recipient: %s\n", info.Recipient)
package main
```

### Testing Standards

#### Table-Driven Tests
```go
func TestExtractRecipient(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
        hasError bool
    }{
        {
            name:     "İş Bankası standard format",
            input:    "ALICI: AHMET YILMAZ\nTUTAR: 1000 TL",
            expected: "AHMET YILMAZ",
            hasError: false,
        },
        {
            name:     "Garanti BBVA format",
            input:    "ALAN: FATMA KAYA\nMİKTAR: 500 TL",
            expected: "FATMA KAYA",
            hasError: false,
        },
        {
            name:     "No recipient found",
            input:    "Some random text without recipient",
            expected: "",
            hasError: true,
        },
        {
            name:     "Turkish characters",
            input:    "ALICI: ÖĞRETMEN ÇAĞLA ŞAHIN",
            expected: "ÖĞRETMEN ÇAĞLA ŞAHIN",
            hasError: false,
        },
    }

    plugin := &Plugin{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := plugin.extractRecipient(tt.input)
            
            if tt.hasError {
                if err == nil {
                    t.Errorf("Expected error but got none")
                }
                return
            }
            
            if err != nil {
                t.Errorf("Unexpected error: %v", err)
                return
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
func BenchmarkPDFProcessing(b *testing.B) {
    plugin := &Plugin{}
    pdfData := loadTestPDF("sample_large.pdf")
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = plugin.processPDF(pdfData)
    }
}

func BenchmarkPatternMatching(b *testing.B) {
    plugin := &Plugin{}
    text := strings.Repeat("ALICI: TEST USER\n", 1000)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = plugin.extractRecipient(text)
    }
}
```

---

## Adding Bank Support

### Step-by-Step Guide

#### 1. Research Bank Format
```go
// Document your findings
/*
Bank: Example Bank
Receipt Types: EFT, Havale, Online Transfer
Field Patterns:
- Recipient: "HESAP SAHİBİ: [NAME]"
- Description: "İŞLEM AÇIKLAMASI: [DESC]"
- Amount: "TOPLAM TUTAR: [AMOUNT] TL"
*/
```

#### 2. Add Pattern Definitions
```go
// Add to existing patterns
var recipientPatterns = []FieldPattern{
    // Existing patterns...
    {
        Name:  "Example Bank Recipient",
        Regex: regexp.MustCompile(`(?i)HESAP SAHİBİ[^:]*:\s*(.+)`),
        Cleaner: cleanRecipientName,
    },
}
```

#### 3. Write Comprehensive Tests
```go
func TestExampleBankExtraction(t *testing.T) {
    tests := []struct {
        name     string
        pdfText  string
        expected TransactionInfo
    }{
        {
            name:    "Example Bank Standard Receipt",
            pdfText: loadTestData("example_bank_receipt.txt"),
            expected: TransactionInfo{
                Recipient:   "JOHN DOE",
                Description: "Payment Description",
                Amount:     "1,000.00 TL",
            },
        },
        {
            name:    "Example Bank EFT Receipt",
            pdfText: loadTestData("example_bank_eft.txt"),
            expected: TransactionInfo{
                Recipient:   "JANE SMITH",
                Description: "EFT Transfer",
                Amount:     "2,500.50 TL",
            },
        },
    }

    plugin := &Plugin{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            info, err := plugin.processPDFText(tt.pdfText)
            if err != nil {
                t.Fatalf("Unexpected error: %v", err)
            }
            
            if info.Recipient != tt.expected.Recipient {
                t.Errorf("Recipient: expected %q, got %q", 
                    tt.expected.Recipient, info.Recipient)
            }
            // Test other fields...
        })
    }
}
```

#### 4. Update Documentation
```markdown
### 🏦 Example Bank

**Status**: ✅ Fully Supported  
**Receipt Types**: EFT, Havale, Online transfers  
**Supported Fields**: Alıcı, Açıklama, İşlem Tutarı  

**Field Patterns**:
- **Recipient**: `HESAP SAHİBİ`, `ALICI`
- **Description**: `İŞLEM AÇIKLAMASI`, `AÇIKLAMA`
- **Amount**: `TOPLAM TUTAR`, `İŞLEM TUTARI`
```

---

## Pull Request Guidelines

### PR Template Checklist

When submitting a PR, ensure you've completed:

- [ ] **Code Quality**
  - [ ] Code follows project style guidelines
  - [ ] No lint warnings or errors
  - [ ] Functions are well-documented
  - [ ] Complex logic is commented

- [ ] **Testing**
  - [ ] New functionality has comprehensive tests
  - [ ] All existing tests pass
  - [ ] Edge cases are covered
  - [ ] Benchmarks added for performance-critical code

- [ ] **Documentation**
  - [ ] README updated if needed
  - [ ] API documentation updated
  - [ ] Changelog entry added
  - [ ] Bank support list updated

- [ ] **Security**
  - [ ] No sensitive data in logs
  - [ ] Input validation implemented
  - [ ] Resource limits respected
  - [ ] Dependencies are secure

### PR Description Template

```markdown
## Description
Brief description of changes made.

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests pass
- [ ] Manual testing completed

## Screenshots (if applicable)
Add screenshots to help explain your changes.

## Checklist
- [ ] My code follows the style guidelines of this project
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
```

---

## Review Process

### What to Expect

1. **Automated Checks** (1-2 minutes):
   - CI/CD pipeline runs
   - Code quality analysis
   - Security scanning
   - Automated testing

2. **Code Review** (1-3 days):
   - Maintainer reviews your code
   - Feedback provided via PR comments
   - Suggestions for improvements

3. **Iteration** (as needed):
   - Address review feedback
   - Make requested changes
   - Re-request review

4. **Approval and Merge**:
   - Code approved by maintainer
   - Automated merge to main branch
   - Release planning (if applicable)

### Review Criteria

**Technical Quality**:
- Code follows established patterns
- Performance considerations addressed
- Error handling is comprehensive
- Tests provide adequate coverage

**Maintainability**:
- Code is readable and well-documented
- Changes are minimal and focused
- No unnecessary complexity introduced
- Future maintenance considered

**Compatibility**:
- Changes don't break existing functionality
- Backward compatibility maintained
- Integration points verified
- Cross-platform considerations

---

## Community Guidelines

### Code of Conduct

We are committed to providing a welcoming and inspiring community for all. Our [Code of Conduct](https://github.com/SkyLostTR/mattermost-dekont-plugin/blob/main/CODE_OF_CONDUCT.md) outlines our expectations for participant behavior.

### Communication

**GitHub Issues**: For bug reports, feature requests, and bank support  
**GitHub Discussions**: For questions, ideas, and community chat  
**Pull Requests**: For code contributions and improvements  
**Email**: For security issues and private communication  

### Recognition

Contributors are recognized in:
- `CONTRIBUTORS.md` file
- Release notes
- GitHub contributor statistics
- Community shout-outs

---

## Getting Help

### Development Questions
- 💬 [GitHub Discussions](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)
- 📖 [Development Guide](development.html)
- 📧 [Email the maintainer](mailto:contact@skylosttr.dev)

### Issue Templates
- 🐛 [Bug Report](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bug_report.md)
- 💡 [Feature Request](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=feature_request.md)
- 🏦 [Bank Support](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bank_support.md)

### Community Resources
- 📚 [Project Wiki](https://github.com/SkyLostTR/mattermost-dekont-plugin/wiki)
- 🎯 [Roadmap](https://github.com/SkyLostTR/mattermost-dekont-plugin/projects)
- 📊 [Project Stats](https://github.com/SkyLostTR/mattermost-dekont-plugin/pulse)

---

## Thank You!

Thank you for contributing to the Mattermost PDF Dekont Parser Plugin! Every contribution, no matter how small, helps make this project better for everyone.

<div class="text-center">
  <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/graphs/contributors" class="btn btn-outline">View All Contributors</a>
</div>
