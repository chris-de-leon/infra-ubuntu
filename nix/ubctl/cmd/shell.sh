#!/usr/bin/env bash

set -eo pipefail

ASSETS_DIR="${1:-}"
if [[ -z "${ASSETS_DIR}" ]]; then
  echo "error: argument 1 (config directory) is required"
  exit 1
fi

CONFIG_DIR="${2:-}"
if [[ -z "${CONFIG_DIR}" ]]; then
  echo "error: argument 2 (assets directory) is required"
  exit 1
fi

NIX_DEV_SH="${CONFIG_DIR}/devsh"

(
  cd "${ASSETS_DIR}" && nix develop --show-trace --no-write-lock-file --profile "${NIX_DEV_SH}" --command bash -c ""
)

nix develop --show-trace --no-write-lock-file "${NIX_DEV_SH}" "${@:2}"
