#!/bin/bash
# Fix permissions script for Mattermost plugin installation
# Run this after extracting the plugin

echo "Setting executable permissions for Linux binary..."
chmod +x plugins/mattermost-dekont-plugin/server/plugin-linux-amd64

echo "Permissions fixed. Plugin should now start correctly."
ls -la plugins/mattermost-dekont-plugin/server/
