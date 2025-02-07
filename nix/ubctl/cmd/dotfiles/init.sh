#!/usr/bin/env bash

set -eo pipefail

# Default value for the -f flag
FORCE_DELETE=false

# Parse command-line options
while getopts "f" opt; do
  case $opt in
  f)
    FORCE_DELETE=true
    ;;
  *)
    echo "Usage: $0 [-f]" >&2
    exit 1
    ;;
  esac
done

# Define configuration directory
CONF_DIR="$HOME/.config"
mkdir -p "$CONF_DIR"

# Create temporary directory and clone the repository
TEMP_DIR="$(mktemp -d)"
trap 'rm -rf $TEMP_DIR' EXIT
git clone 'https://github.com/chris-de-leon/dotfiles' "$TEMP_DIR"

# Function to handle file or directory removal with optional force flag
remove_item() {
  local file_path="$1"

  if [ "$FORCE_DELETE" = false ]; then
    read -rp "Are you sure you want to remove \"$file_path\"? (y/n) " ANSWER
    if [ "$ANSWER" = 'n' ]; then
      echo "Aborting"
      exit 0
    fi
  fi

  rm -rvf "$file_path"
}

# Move files and directories
for item in "starship.toml" "nvim" "tmux"; do
  SRC="$TEMP_DIR/$item"
  DST="$CONF_DIR/$item"

  if [ -e "$DST" ]; then
    remove_item "$DST"
  fi

  mv -v "$SRC" "$DST"
done
