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

# Set executable permissions using WSL/Git Bash and create tar with preserved permissions
Write-Host "Setting executable permissions and creating bundle..."
try {
    # Try using WSL first
    $wslAvailable = $false
    try {
        wsl --status 2>$null | Out-Null
        $wslAvailable = $true
    } catch {
        # WSL not available
    }
    
    if ($wslAvailable) {
        Write-Host "Using WSL to set permissions and create tar..."
        # Convert Windows path to WSL path
        $wslPath = wsl wslpath -a (Get-Location).Path
        wsl chmod +x "$wslPath/dist/server/plugin-linux-amd64"
        wsl tar -czf "$wslPath/$BundleName" -C "$wslPath/dist" .
        Write-Host "Bundle created with WSL (permissions preserved)" -ForegroundColor Green
    } else {
        # Fallback to Git Bash if available
        try {
            bash -c "chmod +x dist/server/plugin-linux-amd64 && tar -czf '$BundleName' -C dist ."
            Write-Host "Bundle created with Git Bash (permissions preserved)" -ForegroundColor Green
        } catch {
            Write-Warning "Could not set executable permissions automatically."
            Write-Warning "Creating bundle without proper permissions - you'll need to run 'chmod +x' on the server."
            # Create tar bundle without preserved permissions as fallback
            Push-Location dist
            tar -czf "..\$BundleName" *
            Pop-Location
        }
    }
} catch {
    Write-Warning "Could not set executable permissions automatically."
    Write-Warning "Creating bundle without proper permissions - you'll need to run 'chmod +x' on the server."
    # Create tar bundle without preserved permissions as fallback
    Push-Location dist
    tar -czf "..\$BundleName" *
    Pop-Location
}

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
