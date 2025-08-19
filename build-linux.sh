#!/bin/bash

# Build script for Mattermost Dekont Plugin (Linux)
set -e

PLUGIN_ID="mattermost-dekont-plugin"
PLUGIN_VERSION=$(grep '"version"' plugin.json | sed 's/.*"version": *"\([^"]*\)".*/\1/')
BUNDLE_NAME="${PLUGIN_ID}-${PLUGIN_VERSION}-linux.tar.gz"

echo "Building Mattermost Dekont Plugin for Linux..."

# Clean previous builds
echo "Cleaning previous builds..."
rm -f plugin *.tar.gz
rm -rf dist/

# Build the plugin for Linux
echo "Building plugin binary..."
GOOS=linux GOARCH=amd64 go build -o plugin

# Create distribution directory
echo "Creating plugin bundle..."
mkdir -p dist/server

# Copy files and set permissions
cp plugin dist/server/plugin-linux-amd64
chmod +x dist/server/plugin-linux-amd64
cp plugin.json dist/

# Create tar bundle
cd dist
tar -czf "../${BUNDLE_NAME}" *
cd ..

echo "Plugin bundle created: ${BUNDLE_NAME}"
echo "Build completed successfully!"

# List the created files
ls -la *.tar.gz plugin

# Clean up the temporary plugin binary
rm -f plugin
