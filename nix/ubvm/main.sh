#!/usr/bin/env bash

set -eo pipefail

SCRIPT_DIR="$(cd -- "$(dirname "$0")" >/dev/null 2>&1 && pwd -P)"
ASSETS_DIR="$(realpath "$SCRIPT_DIR/..")"

op="$1"
case "$op" in
init)
  "$ASSETS_DIR/cmd/init.sh" "$ASSETS_DIR/playbooks"
  ;;
undo)
  "$ASSETS_DIR/cmd/undo.sh" "$ASSETS_DIR/playbooks"
  ;;
*)
  echo "Invalid option: $op"
  echo "Usage: $0 {init|undo}"
  exit 1
  ;;
esac
