#!/usr/bin/env bash

set -eo pipefail

# GitHub repo details
REPO_OWNER="chris-de-leon"
REPO_NAME="infra-ubuntu"

# Infer OS
OS="$(uname -s | tr '[:upper:]' '[:lower:]')" # linux / darwin

# Infer architecture
ARCH="$(uname -m)"
if [[ "${ARCH}" == "arm64" ]] || [[ "${ARCH}" == "aarch64" ]]; then
  ARCH="arm64"
elif [[ "${ARCH}" == "x86_64" ]]; then
  ARCH="amd64"
else
  echo "Unsupported architecture: ${ARCH}"
  exit 1
fi

# Fetch the latest version tag from GitHub API
echo "info: fetching latest version..."
LATEST_VERSION=$(curl -s "https://api.github.com/repos/${REPO_OWNER}/${REPO_NAME}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"tag_name": "v?([^"]+)".*/\1/')
if [[ -z "${LATEST_VERSION}" ]]; then
  echo "error: failed to fetch the latest version."
  exit 1
fi
echo "info: latest version: ${LATEST_VERSION}"

# Construct the download URL
BINARY_NAME="ubctl_${LATEST_VERSION}_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${LATEST_VERSION}/${BINARY_NAME}"

# Create a temporary directory for downloads
TEMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TEMP_DIR"' EXIT

# Download the archive to the temp directory
echo "info: downloading ${BINARY_NAME} to ${TEMP_DIR}..."
curl -s -L -o "${TEMP_DIR}/${BINARY_NAME}" "${DOWNLOAD_URL}"

# Extract the binary
echo "info: extracting ${BINARY_NAME}..."
tar -xzf "${TEMP_DIR}/${BINARY_NAME}" -C "${TEMP_DIR}"

# Move to /usr/local/bin
echo "info: moving ubctl to /usr/local/bin..."
sudo mv "${TEMP_DIR}/ubctl" /usr/local/bin/ubctl

# Grant execution permissions
sudo chmod +x /usr/local/bin/ubctl

# Verify installation
echo "info: verifying installation..."
VERSION="$(ubctl version)"
echo "info: successfully downloaded ubctl version $VERSION"
