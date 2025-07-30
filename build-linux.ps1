# PowerShell Build script for Mattermost Dekont Plugin (Linux target)
param(
    [string]$Version = "1.0.0"
)

$PluginId = "mattermost-dekont-plugin"
$BundleName = "$PluginId-$Version-linux.tar.gz"

Write-Host "Building Mattermost Dekont Plugin for Linux..." -ForegroundColor Green

# Clean previous builds
Write-Host "Cleaning previous builds..."
Remove-Item -Force -ErrorAction SilentlyContinue plugin, *.tar.gz
Remove-Item -Recurse -Force -ErrorAction SilentlyContinue dist/

# Build the plugin for Linux
Write-Host "Building plugin binary..."
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o plugin

if (-not (Test-Path "plugin")) {
    Write-Error "Failed to build plugin binary"
    exit 1
}

# Create distribution directory
Write-Host "Creating plugin bundle..."
New-Item -ItemType Directory -Force -Path "dist\server" | Out-Null

# Copy files
Copy-Item plugin dist\server\plugin-linux-amd64
Copy-Item plugin.json dist\

# Create tar bundle
Push-Location dist
tar -czf "..\$BundleName" *
Pop-Location

if (Test-Path $BundleName) {
    Write-Host "Plugin bundle created: $BundleName" -ForegroundColor Green
    Write-Host "Build completed successfully!" -ForegroundColor Green
    
    # Show file info
    Get-ChildItem $BundleName | Format-Table Name, Length, LastWriteTime
} else {
    Write-Error "Failed to create bundle"
    exit 1
}

# Clean up the temporary plugin binary
Remove-Item -Force plugin

Write-Host "Ready for deployment to Linux Mattermost server!" -ForegroundColor Cyan
