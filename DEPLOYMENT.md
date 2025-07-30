# Deployment Guide for Mattermost PDF Dekont Parser Plugin

## Prerequisites
- Mattermost Server v6.0.0 or higher
- System Administrator privileges in Mattermost

## Installation Steps

1. **Build the Plugin** (if not already done):
   ```bash
   # For Linux deployment (recommended for Mattermost servers)
   $env:GOOS = "linux"; $env:GOARCH = "amd64"; go build -o dist/server/plugin-linux-amd64 .
   
   # For Windows deployment
   go build -o dist/server/plugin-windows-amd64.exe
   ```

2. **Create Plugin Bundle** (if not already done):
   ```bash
   # Create distribution structure
   mkdir -Force dist/server
   copy plugin.json dist/
   
   # Create Linux bundle (recommended)
   cd dist
   tar -czf ../mattermost-dekont-plugin-1.0.0-linux.tar.gz server/plugin-linux-amd64 plugin.json
   cd ..
   ```

3. **Upload to Mattermost**:
   - Go to **System Console** > **Plugins** > **Management**
   - Click **Choose File** and select `mattermost-dekont-plugin-1.0.0-linux.tar.gz`
   - Click **Upload**
   - Wait for the upload to complete

4. **Fix Permissions** (Linux only):
   If you get "permission denied" error, run this on your Mattermost server:
   ```bash
   chmod +x plugins/mattermost-dekont-plugin/server/plugin-linux-amd64
   ```

5. **Enable the Plugin**:
   - In the **System Console** > **Plugins** > **Management**
   - Find "PDF Dekont Parser" in the list
   - Click **Enable**

## Usage

1. Upload a PDF bank receipt to any Mattermost channel
2. The plugin will automatically:
   - Detect the PDF file
   - Extract text content
   - Parse transaction details (Alıcı, Açıklama, İşlem Tutarı)
   - Update the message with formatted information

## Troubleshooting

### Plugin Not Activating
- Check Mattermost server logs for errors
- Ensure Mattermost version is v6.0.0+
- Verify plugin is properly uploaded and enabled

### PDF Processing Issues
- Check that the PDF contains text (not just images)
- Verify the PDF follows Turkish bank receipt format
- Check server logs for specific error messages

### Permission Issues (Linux)
- **Error: "permission denied" when starting plugin**
  - Solution: Run `chmod +x plugins/mattermost-dekont-plugin/server/plugin-linux-amd64` on your server
  - Alternative: Use the included `fix-permissions.sh` script
- Ensure the plugin has write access to temporary directories
- Check that the Mattermost server process has appropriate file permissions

## Configuration

This plugin requires no additional configuration. It automatically activates when enabled and processes all PDF uploads in all channels.

## Support

For issues or questions:
1. Check the Mattermost server logs
2. Verify the PDF format matches expected Turkish bank receipt patterns
3. Ensure all dependencies are properly installed

## Development

To modify or extend the plugin:
1. Edit `plugin.go` for core functionality
2. Update `plugin.json` for metadata changes
3. Rebuild using `go build -o plugin.exe`
4. Create new bundle and re-upload
