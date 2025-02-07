#!/usr/bin/env bash

set -eo pipefail

SCRIPT_DIR="$(cd -- "$(dirname "$0")" >/dev/null 2>&1 && pwd -P)"
ASSETS_DIR="$(realpath "$SCRIPT_DIR/..")"
ansible-playbook "$ASSETS_DIR/playbooks/cli/undo.yml"
