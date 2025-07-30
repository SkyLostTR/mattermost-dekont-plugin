# GitHub Copilot Instructions for Mattermost PDF Dekont Parser Plugin

## Project Context
This is a Mattermost plugin written in Go that automatically parses PDF bank receipts (dekont) and extracts transaction details.

## Key Components
- **Plugin Framework**: Uses Mattermost Server v6 plugin framework
- **PDF Processing**: Uses `github.com/ledongthuc/pdf` for text extraction
- **Field Extraction**: Regex-based pattern matching for Turkish bank receipt fields
- **File Processing**: Handles PDF file uploads in Mattermost channels

## Code Patterns to Follow

### Error Handling
Always log errors using the Mattermost API logger:
```go
p.API.LogError("Description", "error", err.Error())
```

### Resource Management
- Always clean up temporary files using `defer os.Remove(tempFile.Name())`
- Close PDF readers with `defer r.Close()`
- Handle file operations safely

### Regex Patterns
Use case-insensitive regex for Turkish text recognition:
```go
regexp.MustCompile(`(?i)PATTERN`)
```

### Plugin Hooks
- `OnActivate()`: Plugin initialization
- `MessageHasBeenPosted()`: Handle file uploads

## Dependencies
- `github.com/mattermost/mattermost-server/v6` - Core plugin framework
- `github.com/ledongthuc/pdf` - PDF text extraction

## Build Process
- Use `go build -o plugin.exe` for Windows builds
- Plugin manifest is in `plugin.json`
- Create bundles with Makefile targets

## Turkish Language Support
The plugin specifically handles Turkish bank receipt formats with fields like:
- Alıcı (Recipient)
- Açıklama (Description) 
- İşlem Tutarı (Transaction Amount)

When suggesting code changes, maintain compatibility with Turkish text encoding and common bank receipt patterns.
