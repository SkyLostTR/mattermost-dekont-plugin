# Mattermost PDF Dekont Parser Plugin

> **Developed with ❤️ by SkyLostTR (@Keeftraum)**
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
- [Credits & Acknowledgments](#-credits--acknowledgments)
- [License](#license)

<p align="center">
  <img src="assets/icon.svg" alt="PDF Dekont Parser Plugin Icon" width="128" height="128"/>
</p>

<p align="center" dir="auto">
   <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/graphs/contributors">
     <img alt="GitHub Contributors" src="https://img.shields.io/github/contributors/SkyLostTR/mattermost-dekont-plugin" style="max-width: 100%;">
   </a>
   <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/issues">
     <img alt="Issues" src="https://img.shields.io/github/issues/SkyLostTR/mattermost-dekont-plugin?color=0088ff" style="max-width: 100%;">
   </a>
   <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/pulls">
     <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/SkyLostTR/mattermost-dekont-plugin?color=0088ff" style="max-width: 100%;">
   </a>
   <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/stargazers">
     <img alt="GitHub stars" src="https://img.shields.io/github/stars/SkyLostTR/mattermost-dekont-plugin?color=yellow" style="max-width: 100%;">
   </a>
   <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/network/members">
     <img alt="GitHub forks" src="https://img.shields.io/github/forks/SkyLostTR/mattermost-dekont-plugin?color=orange" style="max-width: 100%;">
   </a>
   <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/blob/main/LICENSE">
     <img alt="License" src="https://img.shields.io/github/license/SkyLostTR/mattermost-dekont-plugin?color=blue" style="max-width: 100%;">
   </a>
  <a href="https://golang.org/">
    <img alt="Go Version" src="https://img.shields.io/badge/go-1.21%2B-blue.svg" style="max-width: 100%;">
  </a>
  <a href="https://mattermost.com/">
    <img alt="Mattermost Version" src="https://img.shields.io/badge/mattermost-v6%2B-green.svg" style="max-width: 100%;">
  </a>
    <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/actions/workflows/ci.yml">
      <img alt="CI/CD Pipeline" src="https://github.com/SkyLostTR/mattermost-dekont-plugin/actions/workflows/ci.yml/badge.svg" style="max-width: 100%;">
    </a>
    <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/actions/workflows/code-quality.yml">
      <img alt="Code Quality" src="https://github.com/SkyLostTR/mattermost-dekont-plugin/actions/workflows/code-quality.yml/badge.svg" style="max-width: 100%;">
    </a>
    <a href="https://goreportcard.com/report/github.com/SkyLostTR/mattermost-dekont-plugin">
      <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/SkyLostTR/mattermost-dekont-plugin" style="max-width: 100%;">
    </a>
    <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/watchers">
      <img alt="Watchers" src="https://img.shields.io/github/watchers/SkyLostTR/mattermost-dekont-plugin?color=lightgrey" style="max-width: 100%;">
    </a>
    <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin/releases">
      <img alt="Latest Release" src="https://img.shields.io/github/v/release/SkyLostTR/mattermost-dekont-plugin?sort=semver" style="max-width: 100%;">
    </a>
    <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin">
      <img alt="Last Commit" src="https://img.shields.io/github/last-commit/SkyLostTR/mattermost-dekont-plugin" style="max-width: 100%;">
    </a>
    <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin">
      <img alt="Top Language" src="https://img.shields.io/github/languages/top/SkyLostTR/mattermost-dekont-plugin" style="max-width: 100%;">
    </a>
    <a href="https://github.com/SkyLostTR/mattermost-dekont-plugin">
      <img alt="Repo Size" src="https://img.shields.io/github/repo-size/SkyLostTR/mattermost-dekont-plugin" style="max-width: 100%;">
    </a>
</p>

<!-- Support Buttons -->
<p align="center">
  <a href="https://coff.ee/keeftraum">
    <img src="https://img.shields.io/badge/Buy&nbsp;Me&nbsp;a&nbsp;Coffee-FFDD00?style=for-the-badge&logo=buymeacoffee&logoColor=black" alt="Buy Me a Coffee">
  </a>
  &nbsp;
  <a href="https://github.com/sponsors/SkyLostTR">
    <img src="https://img.shields.io/badge/GitHub&nbsp;Sponsor-ff69b4?style=for-the-badge&logo=github&logoColor=white" alt="GitHub Sponsor">
  </a>
</p>

A Mattermost plugin that automatically parses PDF bank receipts (dekont) and extracts transaction details to display them in a structured format.

## 📚 Documentation

📚 **[Complete Documentation Website](https://skylosttr.github.io/mattermost-dekont-plugin)** - Comprehensive guides, API reference, and examples

Quick Links:
- 🚀 [Installation Guide](https://skylosttr.github.io/mattermost-dekont-plugin/installation.html)
- 📘 [Usage Guide](https://skylosttr.github.io/mattermost-dekont-plugin/usage.html)
- 🏦 [Supported Banks](https://skylosttr.github.io/mattermost-dekont-plugin/banks.html)
- 💻 [Development Guide](https://skylosttr.github.io/mattermost-dekont-plugin/development.html)

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

# Build for Linux (recommended)
$env:GOOS = "linux"; $env:GOARCH = "amd64"; go build -o dist/server/plugin-linux-amd64 .
copy plugin.json dist/
cd dist
tar -czf ../mattermost-dekont-plugin-1.0.0-linux.tar.gz server/plugin-linux-amd64 plugin.json
```

### Installation Steps

1. **Upload Plugin**:
   - Go to **System Console** > **Plugins** > **Management**
   - Click **Choose File** and select `mattermost-dekont-plugin-1.0.0-linux.tar.gz`
   - Click **Upload**

2. **Fix Permissions** (Linux servers only):
   ```bash
   chmod +x plugins/mattermost-dekont-plugin/server/plugin-linux-amd64
   ```

3. **Enable Plugin**:
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

The plugin offers comprehensive customization options accessible through **System Console** > **Plugins** > **PDF Dekont Parser**.

#### 🔧 Core Settings

| Setting | Description | Default | Type |
|---------|-------------|---------|------|
| **Enable PDF Dekont Parser** | Master switch to enable/disable the plugin | `true` | Boolean |
| **Process Only in Specific Channels** | Restrict processing to specified channels | `false` | Boolean |
| **Allowed Channels** | Comma-separated list of channel names | `""` | Text |
| **Maximum File Size (MB)** | PDF size limit for processing | `10` | Number |

#### 🎨 Customization Settings

| Setting | Description | Default | Type |
|---------|-------------|---------|------|
| **Custom Message Prefix** | Text added before extracted data | `📄 **Dekont Bilgileri:**` | Text |
| **Include Processing Timestamp** | Add timestamp to processed messages | `false` | Boolean |
| **Notify on Processing Errors** | Send error messages to channels | `false` | Boolean |
| **Error Notification Message** | Custom error message text | Turkish error message | Text |

#### 🛠️ Advanced Settings

| Setting | Description | Default | Type |
|---------|-------------|---------|------|
| **Enable Debug Logging** | Detailed logging for troubleshooting | `false` | Boolean |
| **Supported Bank Formats** | Read-only list of supported banks | All supported banks | Text |

#### Configuration Examples

**Channel Restriction Example:**
```
Process Only in Specific Channels: ✓ Enabled
Allowed Channels: finance,accounting,treasury,payments
```

**Custom Message Example:**
```
Custom Message Prefix: 💰 Bank Transaction Details:
Include Processing Timestamp: ✓ Enabled
```

**Error Handling Example:**
```
Notify on Processing Errors: ✓ Enabled
Error Notification Message: ⚠️ Unable to process PDF receipt. Please verify it's a valid bank receipt.
```

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
4. 📧 Contact: [keeftraum@protonmail.com]

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

Instead, email us at: [keeftraum@protonmail.com](mailto:keeftraum@protonmail.com)

## 🎉 Credits & Acknowledgments

### 👨‍💻 Development Team

**Primary Developer**: **SkyLostTR (@Keeftraum)** 🚀
- Lead developer and maintainer
- Architecture design and implementation
- Turkish banking integration specialist
- PDF processing optimization

### 🏛️ Banking Partners

We acknowledge the Turkish banking institutions whose receipt formats are supported:
- **Türkiye İş Bankası** - Turkey's leading private bank
- **Garanti BBVA** - Digital banking pioneer
- **Akbank** - Innovation in financial services
- **Yapı Kredi** - Comprehensive banking solutions
- **Ziraat Bankası** - Turkey's largest state bank
- **VakıfBank** - Traditional banking excellence
- **Kuveyt Türk** - Islamic banking leader
- **HalkBank** - People's bank

### 🛠️ Technology Stack

- **[Go](https://golang.org/)** - High-performance backend language
- **[Mattermost Plugin Framework](https://developers.mattermost.com/extend/plugins/)** - Robust plugin architecture
- **[ledongthuc/pdf](https://github.com/ledongthuc/pdf)** - PDF text extraction library
- **[GitHub Actions](https://github.com/features/actions)** - CI/CD pipeline
- **[SonarCloud](https://sonarcloud.io/)** - Code quality analysis

### 🌟 Special Thanks

- **Mattermost Community** for the excellent plugin framework
- **Turkish Banking Association** for standardized receipt formats
- **Open Source Contributors** who make projects like this possible
- **Beta Testers** who helped refine the bank format recognition

### 💼 Professional Support

For enterprise implementations, custom bank format support, or professional services:

**Contact SkyLostTR (@Keeftraum)** 📧
- Email: [keeftraum@protonmail.com](mailto:keeftraum@protonmail.com)
- GitHub: [@SkyLostTR](https://github.com/SkyLostTR)

---

*"Built with ❤️ for the Mattermost community by SkyLostTR (@Keeftraum)"*

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### License Summary

- ✅ **Commercial Use** - Use in commercial projects
- ✅ **Modification** - Modify and adapt the code
- ✅ **Distribution** - Distribute original or modified versions
- ✅ **Private Use** - Use privately in your organization
- ❗ **Include License** - Include the original license in distributions
- ❗ **Include Copyright** - Include original copyright notice

**Developed by SkyLostTR (@Keeftraum)** - Contributing to the open source community

## 📞 Contact

For questions, support, or collaboration:
- **Email**: [keeftraum@protonmail.com](mailto:keeftraum@protonmail.com)
- **GitHub**: [@SkyLostTR](https://github.com/SkyLostTR)
- **Issues**: [Project Issues](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues)
- **Discussions**: [GitHub Discussions](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)

---

<p align="center">
  <strong>⭐ If this project helped you, please give it a star! ⭐</strong><br>
  <em>Developed with passion by SkyLostTR (@Keeftraum) for the global Mattermost community</em>
</p>

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
