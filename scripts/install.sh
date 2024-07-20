#!/bin/bash

# Exit on any error
set -e

REPO="OpenStruct/goeasycli"

fetch_latest_tag() {
    local repo_url="$1"
    local api_url="https://api.github.com/repos/${repo_url#https://github.com/}/tags"
    
    if command -v curl >/dev/null 2>&1; then
        CLI_VERSION=$(curl -sSL "$api_url" | grep -o '"name": "[^"]*' | sed 's/"name": "//g' | head -n 1)
    elif command -v wget >/dev/null 2>&1; then
        CLI_VERSION=$(wget -qO- "$api_url" | grep -o '"name": "[^"]*' | sed 's/"name": "//g' | head -n 1)
    else
        echo "Error: Neither curl nor wget is installed. Cannot fetch latest tag." >&2
        exit 1
    fi

    if [ -z "$CLI_VERSION" ]; then
        echo "Error: No tags found in the repository." >&2
        exit 1
    else
        echo "Latest tag found: $CLI_VERSION"
    fi
}

# Try to use git if available
if command -v git >/dev/null 2>&1; then
    CLI_VERSION=$(git ls-remote --tags --refs --sort="version:refname" --exit-code \
      https://github.com/$REPO.git | tail -n1 | sed 's/.*\///')
    
    if [ -z "$CLI_VERSION" ]; then
        echo "Git command failed. Falling back to API method."
        fetch_latest_tag "$REPO"
    else
        echo "Latest tag found (via git): $CLI_VERSION"
    fi
else
    echo "Git is not installed. Using API method to fetch latest tag."
    fetch_latest_tag "$REPO"
fi

# BASE_URL="file:///Users/noahalorwu/Desktop/goeasycli/dist/"
BASE_URL="https://github.com/$REPO/releases/download/$CLI_VERSION"

# Determine the OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $OS in
  linux)
    OS="linux"
    ;;
  darwin)
    OS="darwin"
    ;;
  msys*|mingw*|cygwin*)
    OS="windows"
    ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

case $ARCH in
  x86_64)
    ARCH="amd64"
    ;;
  arm64|aarch64)
    ARCH="arm64"
    ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

# Construct the download URL
URL="$BASE_URL/goeasycli_${OS}_${ARCH}.tar.gz"

# Download and extract the binary
TEMP_DIR=$(mktemp -d)
# chmod u+rwx "$TEMP_DIR"
echo "Downloading $URL..."
curl -L -o "$TEMP_DIR/goeasycli.tar.gz" "$URL"
tar -xvzf "$TEMP_DIR/goeasycli.tar.gz" -C "$TEMP_DIR"

# Move the binary to /usr/local/bin (or another directory in PATH)
sudo mv "$TEMP_DIR/goeasycli" /usr/local/bin/goeasycli
sudo chmod +x /usr/local/bin/goeasycli

# Clean up
rm -rf "$TEMP_DIR"

# Verify installation
if command -v goeasycli &> /dev/null
then
    echo "goeasycli installed successfully!"
else
    echo "Installation failed."
    exit 1
fi
