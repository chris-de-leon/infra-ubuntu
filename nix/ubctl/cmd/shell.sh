#!/usr/bin/env bash

set -eo pipefail

NIX_URL="$1"
if [ -z "$NIX_URL" ]; then
  echo "error: argument 1 (nix url) is required"
  exit 1
fi

nix develop "$NIX_URL"
