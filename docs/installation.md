---
layout: default
title: Installation Guide
nav_order: 2
description: "Step-by-step installation guide for Mattermost PDF Dekont Parser Plugin"
---

# Installation Guide
{: .no_toc }

Complete guide to installing and configuring the Mattermost PDF Dekont Parser Plugin.
{: .fs-6 .fw-300 }

## Table of contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Prerequisites

Before installing the plugin, ensure you have:

- **Mattermost Server v6.0.0 or higher**
- **System Administrator privileges** in Mattermost
- **Plugin uploads enabled** in System Console

### Compatibility Matrix

| Mattermost Version | Plugin Version | Status |
|:-------------------|:---------------|:-------|
| v6.0.0+ | v1.0.0+ | ✅ Supported |
| v5.x | - | ❌ Not supported |

---

## Installation Methods

### Method 1: Pre-built Release (Recommended)

This is the easiest way to install the plugin.

#### Step 1: Download
1. Go to the [Releases page](https://github.com/SkyLostTR/mattermost-dekont-plugin/releases)
2. Download the latest `mattermost-dekont-plugin-X.X.X.tar.gz` file

#### Step 2: Upload to Mattermost
1. Navigate to **System Console** → **Plugins** → **Management**
2. Click **Choose File** and select the downloaded `.tar.gz` file
3. Click **Upload**
4. Wait for the upload progress to complete

#### Step 3: Enable Plugin
1. Find "PDF Dekont Parser" in the installed plugins list
2. Click **Enable** button
3. Verify the plugin status shows as "Running"

### Method 2: Build from Source

For developers or advanced users who want to build from source.

#### Step 1: Clone Repository
```bash
git clone https://github.com/SkyLostTR/mattermost-dekont-plugin.git
cd mattermost-dekont-plugin
```

#### Step 2: Install Dependencies
```bash
go mod download
```

#### Step 3: Build Plugin
```bash
# Using Make (recommended)
make build
make bundle

# Or manually
go build -o plugin.exe
mkdir -p dist/server
cp plugin.exe dist/server/plugin-windows-amd64.exe
cp plugin.json dist/
cd dist && tar -czf ../mattermost-dekont-plugin-1.0.0.tar.gz *
```

#### Step 4: Install Built Plugin
Follow steps 2-3 from Method 1 using your built `.tar.gz` file.

---

## Configuration

### System Console Settings

After installation, you can configure the plugin:

1. Go to **System Console** → **Plugins** → **PDF Dekont Parser**
2. Configure available settings (currently minimal configuration required)

### Plugin Settings

| Setting | Default | Description |
|:--------|:--------|:------------|
| Enable Plugin | `false` | Enable/disable the plugin |
| Debug Logging | `false` | Enable detailed debug logs |

---

## Verification

### Test Installation

1. **Upload a Test PDF**:
   - Go to any Mattermost channel
   - Upload a PDF bank receipt
   - Check if the plugin processes it automatically

2. **Check Server Logs**:
   ```
   Look for entries like:
   [INFO] PDF Dekont Parser: Plugin activated successfully
   [INFO] PDF Dekont Parser: Processing PDF file: filename.pdf
   ```

3. **Verify Plugin Status**:
   - **System Console** → **Plugins** → **Management**
   - Plugin should show status "Running" with green indicator

---

## Troubleshooting

### Common Issues

#### Plugin Not Activating

**Symptoms**: Plugin shows as "Failed" or doesn't start

**Solutions**:
1. Check Mattermost server logs:
   ```bash
   tail -f /opt/mattermost/logs/mattermost.log
   ```
2. Verify Mattermost version compatibility
3. Ensure plugin uploads are enabled in System Console
4. Try re-uploading the plugin

#### PDF Files Not Being Processed

**Symptoms**: PDFs upload but no extraction occurs

**Solutions**:
1. Check file format (must be PDF)
2. Verify the PDF contains text (not scanned images)
3. Check server logs for processing errors
4. Ensure the PDF is from a supported bank

#### Missing Transaction Fields

**Symptoms**: Some fields not extracted from PDF

**Solutions**:
1. Check if your bank format is supported
2. [Submit a bank support request](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=bank_support.md)
3. Enable debug logging for detailed extraction info

### Error Messages

#### "Plugin failed to start"
- **Cause**: Usually a configuration or dependency issue
- **Fix**: Check system requirements and re-upload plugin

#### "PDF processing failed"
- **Cause**: Malformed PDF or unsupported format
- **Fix**: Try with different PDF or check bank support

#### "Permission denied"
- **Cause**: Insufficient file system permissions
- **Fix**: Check Mattermost server file permissions

---

## Uninstallation

To remove the plugin:

1. Go to **System Console** → **Plugins** → **Management**
2. Find "PDF Dekont Parser" in the list
3. Click **Disable** to stop the plugin
4. Click **Remove** to uninstall completely

---

## Next Steps

After successful installation:

1. [Learn how to use the plugin](usage.html)
2. [Check supported banks](banks.html)
3. [Explore development options](development.html)

---

## Support

Need help with installation?

- 📖 [Usage Guide](usage.html)
- 🐛 [Report Installation Issues](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues)
- 💬 [Community Discussions](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)
- 📧 [Contact Developer](mailto:keeftraum@protonmail.com)
