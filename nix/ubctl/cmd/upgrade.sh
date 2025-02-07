#!/usr/bin/env bash

set -eo pipefail

PLAYBOOKS_DIR="$1"
if [ -z "$PLAYBOOKS_DIR" ]; then
  echo "error: argument 1 (playbooks directory) is required"
  exit 1
fi

ansible-playbook "$PLAYBOOKS_DIR/bashrc/init.version.yml"
