# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **🚀 Comprehensive Plugin Settings** - Full customization options for administrators
  - Enable/disable plugin functionality
  - Channel-specific processing controls
  - File size limits configuration
  - Custom message prefixes and formatting
  - Processing timestamp options
  - Error notification settings
  - Debug logging controls
- **🎨 Enhanced User Experience**
  - Professional plugin icon designed by SkyLostTR (@Keeftraum)
  - Improved message formatting with credits footer
  - Configurable error messages and notifications
  - Real-time processing feedback
- **⚙️ Advanced Configuration System**
  - Complete settings schema for Mattermost System Console
  - Proper configuration validation and defaults
  - Channel allow-list functionality
  - File size restrictions for better performance
- **📚 Enhanced Documentation**
  - Comprehensive README with plugin settings guide
  - Detailed configuration examples and use cases
  - Professional credits and acknowledgments section
  - Technology stack recognition
- **🏗️ Code Architecture Improvements**
  - Configuration management system
  - Proper error handling and logging
  - Enhanced debugging capabilities
  - Credit attribution throughout codebase
- **💼 Professional Features**
  - Plugin marketplace information (author, homepage, support URLs)
  - Professional icon and branding
  - Comprehensive help text and documentation
  - SkyLostTR (@Keeftraum) brand integration
- Comprehensive test suite with unit tests and benchmarks
- GitHub Actions CI/CD pipeline with multi-platform builds
- Security scanning with Gosec and Trivy
- Code quality analysis with golangci-lint and SonarCloud
- Automated dependency updates
- Issue and PR templates
- Contributing guidelines
- Code coverage reporting

### Changed
- Improved error handling and logging
- Enhanced PDF parsing with better regex patterns
- Updated documentation with detailed setup instructions

### Security
- Added security scanning to CI pipeline
- Implemented input validation best practices

## [1.0.0] - 2025-07-30

### Added
- Initial release of Mattermost PDF Dekont Parser Plugin
- PDF text extraction using github.com/ledongthuc/pdf
- Turkish bank receipt field recognition
- Support for common Turkish bank formats
- Automatic message updating with extracted transaction details
- Regex-based parsing for:
  - Alıcı (Recipient) field
  - Açıklama (Description) field  
  - İşlem Tutarı (Transaction Amount) field
- Error handling and logging via Mattermost API
- Resource cleanup for temporary files
- Plugin manifest for Mattermost compatibility

### Features
- **Automatic Detection**: Monitors PDF file uploads in all channels
- **Multi-format Support**: Handles various Turkish bank receipt formats
- **Real-time Processing**: Immediate extraction and display of transaction data
- **Error Resilience**: Graceful handling of malformed or unsupported PDFs
- **Turkish Language Support**: Full support for Turkish characters and banking terminology

### Technical Details
- Built with Go and Mattermost Server v6 plugin framework
- Uses case-insensitive regex patterns for robust field matching
- Implements proper resource management with defer statements
- Optimized for Linux deployment
