# Deployment Guide for Mattermost PDF Dekont Parser Plugin

## Prerequisites
- Mattermost Server v6.0.0 or higher
- System Administrator privileges in Mattermost

## Installation Steps

1. **Build the Plugin** (if not already done):
   ```bash
   go build -o plugin.exe
   ```

2. **Create Plugin Bundle** (if not already done):
   ```bash
   # Create distribution structure
   mkdir -Force dist/server
   copy plugin.exe dist/server/plugin-windows-amd64.exe
   copy plugin.json dist/
   
   # Create bundle
   cd dist
   tar -czf ../mattermost-dekont-plugin-1.0.0.tar.gz *
   cd ..
   ```

3. **Upload to Mattermost**:
   - Go to **System Console** > **Plugins** > **Management**
   - Click **Choose File** and select `mattermost-dekont-plugin-1.0.0.tar.gz`
   - Click **Upload**
   - Wait for the upload to complete

4. **Enable the Plugin**:
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

### Permission Issues
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
