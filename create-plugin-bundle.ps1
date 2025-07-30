# Create Linux plugin bundle with correct permissions
$bundleName = "mattermost-dekont-plugin-1.0.0-linux.tar.gz"

# Remove existing bundle if it exists
if (Test-Path $bundleName) {
    Remove-Item $bundleName
}

# Create tar with executable permissions for Linux binary
# We'll use a workaround by creating the tar and then using WSL/Git Bash if available
if (Get-Command wsl -ErrorAction SilentlyContinue) {
    Write-Host "Using WSL to create bundle with correct permissions..."
    wsl tar --mode='a+x' -czf $bundleName -C dist server/plugin-linux-amd64 plugin.json
} elseif (Get-Command bash -ErrorAction SilentlyContinue) {
    Write-Host "Using Git Bash to create bundle with correct permissions..."
    bash -c "tar --mode='a+x' -czf $bundleName -C dist server/plugin-linux-amd64 plugin.json"
} else {
    Write-Host "Creating bundle with Windows tar (manual permission fix may be needed)..."
    tar -czf $bundleName -C dist server/plugin-linux-amd64 plugin.json
    Write-Host "Warning: You may need to manually set executable permissions on the Linux binary after extraction."
}

Write-Host "Bundle created: $bundleName"
