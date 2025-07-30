# Mattermost PDF Dekont Parser Plugin

A Mattermost plugin that automatically parses PDF bank receipts (dekont) and extracts transaction details to display them in a structured format.

## Features

- **Automatic PDF Processing**: Detects PDF files uploaded to Mattermost channels
- **Text Extraction**: Extracts text content from PDF documents
- **Field Recognition**: Identifies key transaction fields:
  - Alıcı (Recipient)
  - Açıklama (Description)
  - İşlem Tutarı (Transaction Amount)
- **Auto-posting**: Updates the message with extracted information

## How it Works

1. When a PDF file is uploaded to any channel
2. The plugin automatically processes the file
3. Extracts transaction details using regex patterns
4. Updates the original post with formatted transaction information

## Installation

1. Build the plugin:
   ```bash
   go build -o plugin.exe
   ```

2. Create a plugin bundle by compressing the plugin files
3. Upload the bundle through Mattermost System Console > Plugins > Management
4. Enable the plugin

## Development

### Prerequisites
- Go 1.19+
- Mattermost Server v6.0.0+

### Building
```bash
go mod tidy
go build -o plugin.exe
```

### Dependencies
- `github.com/mattermost/mattermost-server/v6` - Mattermost plugin framework
- `github.com/ledongthuc/pdf` - PDF text extraction

## Configuration

No additional configuration is required. The plugin automatically activates when enabled.

## Supported PDF Formats

The plugin works with Turkish bank receipt PDFs and recognizes common field patterns used by Turkish banks.

## License

MIT License
