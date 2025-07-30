# Mattermost PDF Dekont Parser Plugin

[![CI/CD Pipeline](https://github.com/SkyLostTR/mattermost-dekont-plugin/actions/workflows/ci.yml/badge.svg)](https://github.com/SkyLostTR/mattermost-dekont-plugin/actions/workflows/ci.yml)
[![Code Quality](https://github.com/SkyLostTR/mattermost-dekont-plugin/actions/workflows/code-quality.yml/badge.svg)](https://github.com/SkyLostTR/mattermost-dekont-plugin/actions/workflows/code-quality.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/SkyLostTR/mattermost-dekont-plugin)](https://goreportcard.com/report/github.com/SkyLostTR/mattermost-dekont-plugin)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/v/release/SkyLostTR/mattermost-dekont-plugin)](https://github.com/SkyLostTR/mattermost-dekont-plugin/releases)

A Mattermost plugin that automatically parses PDF bank receipts (dekont) and extracts transaction details to display them in a structured format.

## 📋 Table of Contents

- [Features](#features)
- [Supported Banks](#supported-banks)
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [Security](#security)
- [License](#license)

## ✨ Features

- **🔍 Automatic PDF Detection**: Monitors PDF file uploads in all Mattermost channels
- **📄 Text Extraction**: Robust PDF text extraction using advanced parsing techniques
- **🏦 Multi-Bank Support**: Handles various Turkish bank receipt formats
- **🇹🇷 Turkish Language Support**: Full support for Turkish characters and banking terminology
- **⚡ Real-time Processing**: Immediate extraction and display of transaction data
- **🛡️ Error Resilience**: Graceful handling of malformed or unsupported PDFs
- **📊 Field Recognition**: Intelligent extraction of key transaction fields:
  - **Alıcı** (Recipient/Beneficiary)
  - **Açıklama** (Description/Reference)
  - **İşlem Tutarı** (Transaction Amount)
- **🔄 Auto-formatting**: Updates posts with structured transaction information
- **📝 Comprehensive Logging**: Detailed error tracking and debugging information

## 🏦 Supported Banks

Currently supports PDF receipt formats from major Turkish banks:

| Bank | Status | Notes |
|------|--------|-------|
| Türkiye İş Bankası | ✅ Supported | EFT, Havale receipts |
| Garanti BBVA | ✅ Supported | Standard transfer receipts |
| Akbank | ✅ Supported | Online banking receipts |
| Yapı Kredi | ✅ Supported | Transfer confirmations |
| Ziraat Bankası | ✅ Supported | Government bank receipts |

> **Note**: If your bank is not listed, please [request support](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bank_support.md) by providing a sample receipt.

## 🚀 Installation

### Prerequisites
- Mattermost Server v6.0.0 or higher
- System Administrator privileges in Mattermost

### Option 1: Download Pre-built Release

1. Go to the [Releases page](https://github.com/SkyLostTR/mattermost-dekont-plugin/releases)
2. Download the latest `mattermost-dekont-plugin-X.X.X.tar.gz`
3. Upload via **System Console** > **Plugins** > **Management**

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/SkyLostTR/mattermost-dekont-plugin.git
cd mattermost-dekont-plugin

# Install dependencies
go mod download

# Build the plugin
make build

# Create plugin bundle
make bundle
```

### Installation Steps

1. **Upload Plugin**:
   - Go to **System Console** > **Plugins** > **Management**
   - Click **Choose File** and select the `.tar.gz` file
   - Click **Upload**

2. **Enable Plugin**:
   - Find "PDF Dekont Parser" in the plugin list
   - Click **Enable**
   - Verify activation in server logs

## 📖 Usage

### Basic Usage

1. **Upload a PDF**: Attach a PDF bank receipt to any Mattermost channel
2. **Automatic Processing**: The plugin detects and processes the PDF automatically
3. **View Results**: Transaction details appear in the same message thread

### Example Output

When you upload a bank receipt PDF, the plugin will update your message with:

```
**Açıklama**: Invoice Payment - INV-2023-001
**Alıcı**: ABC Company Ltd.
**İşlem Tutarı**: 1,500.00 TL
```

### Supported PDF Types

- ✅ EFT (Electronic Funds Transfer) receipts
- ✅ Wire transfer confirmations
- ✅ Online banking transaction receipts
- ✅ ATM transfer receipts
- ✅ Mobile banking confirmations

## 🛠️ Development

### Setup Development Environment

```bash
# Prerequisites
go version  # Requires Go 1.21+
git --version

# Clone and setup
git clone https://github.com/SkyLostTR/mattermost-dekont-plugin.git
cd mattermost-dekont-plugin
go mod download
```

### Available Commands

```bash
# Build plugin
make build

# Run tests
make test

# Create plugin bundle
make bundle

# Clean build artifacts
make clean

# Format code
make fmt

# Lint code
make lint
```

### VS Code Integration

The project includes VS Code tasks for seamless development:

- **Ctrl+Shift+P** → "Tasks: Run Task"
- Choose from: Build, Create Bundle, Clean, Install Dependencies

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=.

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Project Structure

```
mattermost-dekont-plugin/
├── plugin.go              # Main plugin logic
├── plugin_test.go          # Comprehensive test suite
├── plugin.json             # Plugin manifest
├── go.mod                  # Go module definition
├── Makefile               # Build automation
├── .github/               # GitHub workflows and templates
│   ├── workflows/         # CI/CD pipelines
│   └── ISSUE_TEMPLATE/    # Issue templates
├── docs/                  # Additional documentation
└── dist/                  # Build output directory
```

## ⚙️ Configuration

### Plugin Settings

The plugin works out-of-the-box with no additional configuration required. Advanced settings may be added in future versions.

### Mattermost Requirements

- **File Upload Limits**: Ensure Mattermost allows PDF uploads
- **Plugin Permissions**: System admin access for installation
- **Server Resources**: Sufficient memory for PDF processing

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `MM_PLUGIN_DEBUG` | Enable debug logging | `false` |

## 🔧 Troubleshooting

### Common Issues

#### Plugin Not Activating
```bash
# Check Mattermost logs
sudo journalctl -u mattermost -f

# Verify plugin installation
ls -la /opt/mattermost/plugins/
```

#### PDF Processing Fails
- ✅ Verify PDF contains text (not just images)
- ✅ Check PDF is from a supported bank
- ✅ Ensure PDF is not password-protected
- ✅ Review server logs for specific errors

#### Performance Issues
- Monitor memory usage during PDF processing
- Check for concurrent file uploads
- Verify adequate server resources

### Debug Mode

Enable debug logging in `config.json`:

```json
{
  "PluginSettings": {
    "Plugins": {
      "mattermost-dekont-plugin": {
        "debug": true
      }
    }
  }
}
```

### Getting Help

1. 📖 Check our [Documentation](https://github.com/SkyLostTR/mattermost-dekont-plugin/wiki)
2. 🐛 [Report Issues](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues)
3. 💬 [Discussions](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)
4. 📧 Contact: [support@example.com]

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Quick Start for Contributors

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

### Development Guidelines

- Follow Go best practices
- Write tests for new features
- Update documentation
- Use conventional commit messages
- Ensure CI passes

## 🔒 Security

Security is a top priority. Please review our [Security Policy](SECURITY.md) for:

- Vulnerability reporting process
- Security best practices
- Supported versions
- Security features

### Reporting Security Issues

🚨 **Do not report security vulnerabilities through public issues**

Instead, email us at: [security@example.com]

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🏆 Acknowledgments

- [Mattermost](https://mattermost.com/) for the excellent plugin framework
- [ledongthuc/pdf](https://github.com/ledongthuc/pdf) for PDF processing capabilities
- Turkish banking community for format specifications
- All contributors and testers

## 📊 Project Status

- ✅ **Stable**: Production ready
- 🔄 **Actively Maintained**: Regular updates and bug fixes
- 📈 **Growing**: New bank formats added regularly
- 🌍 **Community Driven**: Open to contributions

---

<div align="center">

**⭐ Star this project if you find it useful!**

[Report Bug](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues) · [Request Feature](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues) · [Documentation](https://github.com/SkyLostTR/mattermost-dekont-plugin/wiki)

</div>
