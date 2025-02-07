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
echo ""

# Function to remove files with optional force flag
remove_file() {
  local file_path="$1"

  echo "Removing \"$file_path\"..."
  if [ "$FORCE_DELETE" = true ]; then
    rm -fvr "$file_path"
  else
    rm -Ivr "$file_path"
  fi
  echo ""
}

# Remove files and directories
remove_file "$CONF_DIR/starship.toml"
remove_file "$CONF_DIR/tmux"
remove_file "$CONF_DIR/nvim"
