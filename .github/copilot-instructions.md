# GitHub Copilot Instructions for Mattermost PDF Dekont Parser Plugin

## Project Context
This is a production-ready Mattermost plugin written in Go that automatically parses PDF bank receipts (dekont) and extracts transaction details. The project follows enterprise software development practices with comprehensive testing, CI/CD, and documentation.

## Key Components
- **Plugin Framework**: Uses Mattermost Server v6 plugin framework
- **PDF Processing**: Uses `github.com/ledongthuc/pdf` for text extraction
- **Field Extraction**: Regex-based pattern matching for Turkish bank receipt fields
- **File Processing**: Handles PDF file uploads in Mattermost channels
- **Testing**: Comprehensive unit tests with benchmarks and edge case coverage
- **CI/CD**: GitHub Actions workflows for testing, building, and security scanning

## Development Workflow

### Code Standards
- Follow Go best practices and idiomatic patterns
- Use `gofmt` for consistent formatting
- Write comprehensive tests for all new functionality
- Maintain test coverage above 80% for core business logic
- Use conventional commit messages (feat:, fix:, docs:, etc.)

### Error Handling
Always log errors using the Mattermost API logger:
```go
p.API.LogError("Description", "error", err.Error())
```

### Resource Management
- Always clean up temporary files using `defer os.Remove(tempFile.Name())`
- Close PDF readers with proper error handling
- Handle file operations safely with appropriate error checks

### Testing Patterns
Use table-driven tests for comprehensive coverage:
```go
tests := []struct {
    name     string
    input    string
    expected string
}{
    {
        name:     "descriptive test name",
        input:    "test input",
        expected: "expected output",
    },
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        result := functionUnderTest(tt.input)
        if result != tt.expected {
            t.Errorf("got %v, want %v", result, tt.expected)
        }
    })
}
```

### Regex Patterns
Use case-insensitive regex for Turkish text recognition:
```go
regexp.MustCompile(`(?i)PATTERN`)
```

### Plugin Hooks
- `OnActivate()`: Plugin initialization and setup
- `MessageHasBeenPosted()`: Handle file uploads and PDF processing

## Project Structure

### Core Files
- `plugin.go`: Main plugin logic and PDF processing
- `plugin_test.go`: Comprehensive test suite with benchmarks
- `plugin.json`: Plugin manifest and configuration
- `go.mod/go.sum`: Go module dependencies

### GitHub Integration
- `.github/workflows/`: CI/CD pipelines for testing, building, security
- `.github/ISSUE_TEMPLATE/`: Bug reports, feature requests, bank support
- `.github/pull_request_template.md`: PR template with checklists
- `CONTRIBUTING.md`: Development guidelines and processes
- `SECURITY.md`: Security policy and vulnerability reporting

### Build and Quality
- `Makefile`: Build automation and common tasks
- `.golangci.yml`: Linting configuration
- `.editorconfig`: Editor consistency settings
- `sonar-project.properties`: Code quality analysis

## Dependencies

### Core Dependencies
- `github.com/mattermost/mattermost-server/v6` - Plugin framework
- `github.com/ledongthuc/pdf` - PDF text extraction

### Development Tools
- golangci-lint for code quality
- GitHub Actions for CI/CD
- SonarCloud for code analysis
- Trivy for security scanning

## Build Process
- Use `make build` for local development
- Use `make bundle` to create deployment packages
- GitHub Actions automatically builds for multiple platforms
- Release process creates tagged bundles

## Turkish Language Support
The plugin specifically handles Turkish bank receipt formats with fields like:
- Alıcı (Recipient) - matches "ALICI", "ALICI AD SOYAD/UNVAN"
- Açıklama (Description) - matches "AÇIKLAMA", "ACIKLAMA"
- İşlem Tutarı (Transaction Amount) - matches various amount field patterns

### Supported Banks
- Türkiye İş Bankası
- Garanti BBVA  
- Akbank
- Yapı Kredi
- Ziraat Bankası

## Performance Considerations
- PDF processing should complete in under 2 seconds for typical files
- Memory usage should remain reasonable for concurrent uploads
- Benchmark tests verify performance characteristics
- Large PDF handling is optimized

## Security Guidelines
- Never log sensitive financial information
- Validate all PDF inputs before processing
- Clean up temporary files immediately after use
- Follow secure coding practices for file handling
- Regular dependency updates via automated PRs

## Documentation Standards
- Update README.md for user-facing changes
- Maintain CHANGELOG.md with semantic versioning
- Include inline comments for complex regex patterns
- Document any new bank format support
- Keep deployment guides current

## When Adding New Bank Support
1. Create issue using bank support template
2. Analyze PDF format and field patterns
3. Add regex patterns with case-insensitive matching
4. Write comprehensive tests for the new format
5. Update documentation with supported bank list
6. Consider edge cases and error handling

## Code Review Checklist
- [ ] Tests added for new functionality
- [ ] Documentation updated
- [ ] Security considerations addressed
- [ ] Performance impact assessed
- [ ] Error handling implemented
- [ ] Turkish character encoding handled
- [ ] Bank format compatibility verified

When suggesting code changes, maintain compatibility with Turkish text encoding and existing bank receipt patterns while following the established testing and documentation standards.
