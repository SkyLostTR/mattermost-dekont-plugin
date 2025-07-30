---
layout: default
title: Home
nav_order: 1
description: "Mattermost PDF Dekont Parser Plugin - Automatically parse PDF bank receipts and extract transaction details"
permalink: /
---

# Mattermost PDF Dekont Parser Plugin
{: .fs-9 }

A powerful Mattermost plugin that automatically parses PDF bank receipts (dekont) and extracts transaction details to display them in a structured format.
{: .fs-6 .fw-300 }

[Get Started Now](installation.html){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 } [View on GitHub](https://github.com/SkyLostTR/mattermost-dekont-plugin){: .btn .fs-5 .mb-4 .mb-md-0 }

---

## Quick Features Overview

### 🔍 Automatic PDF Detection
Monitors PDF file uploads in all Mattermost channels and automatically processes bank receipts.

### 🏦 Multi-Bank Support
Supports major Turkish banks including İş Bankası, Garanti BBVA, Akbank, Yapı Kredi, and Ziraat Bankası.

### 🇹🇷 Turkish Language Support
Full support for Turkish characters and banking terminology with intelligent field recognition.

### ⚡ Real-time Processing
Immediate extraction and display of transaction data with structured formatting.

---

## What Gets Extracted?

The plugin intelligently identifies and extracts key transaction fields:

- **Alıcı** (Recipient/Beneficiary)
- **Açıklama** (Description/Reference) 
- **İşlem Tutarı** (Transaction Amount)

## Example Output

When you upload a PDF bank receipt, the plugin automatically updates your message with structured information:

```
📄 PDF Dekont Bilgileri:
━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 Alıcı: JOHN DOE
📝 Açıklama: Freelance Payment
💰 İşlem Tutarı: 1,500.00 TL

━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

## Why Choose This Plugin?

### Enterprise-Ready
- Production-tested codebase
- Comprehensive error handling
- Detailed logging and monitoring
- CI/CD pipelines with automated testing

### Developer-Friendly
- Open source with MIT license
- Well-documented API
- Extensive test coverage (>80%)
- Active community support

### Security-Focused
- No sensitive data logging
- Secure file handling
- Regular security updates
- Vulnerability scanning

---

## Quick Start

1. **Download** the latest release from GitHub
2. **Upload** to your Mattermost server via System Console
3. **Enable** the plugin in Plugin Management
4. **Upload** a PDF bank receipt to test

[Detailed Installation Guide →](installation.html)

---

## Community & Support

- 🐛 [Report Issues](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues)
- 💡 [Request Features](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=feature_request.md)
- 🏦 [Request Bank Support](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bank_support.md)
- 🤝 [Contributing Guide](contributing.html)

---

## Latest Release

[![GitHub release](https://img.shields.io/github/v/release/SkyLostTR/mattermost-dekont-plugin?sort=semver)](https://github.com/SkyLostTR/mattermost-dekont-plugin/releases)
[![Downloads](https://img.shields.io/github/downloads/SkyLostTR/mattermost-dekont-plugin/total)](https://github.com/SkyLostTR/mattermost-dekont-plugin/releases)

Check out the [changelog](changelog.html) for what's new in the latest version.

---

<div class="text-center">
  <a href="https://coff.ee/keeftraum" class="btn btn-outline">☕ Buy Me a Coffee</a>
  <a href="https://github.com/sponsors/SkyLostTR" class="btn btn-outline">💖 GitHub Sponsor</a>
</div>
