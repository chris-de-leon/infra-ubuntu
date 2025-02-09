#!/usr/bin/env bash

set -eo pipefail

PLAYBOOKS_DIR="${1:-}"
if [[ -z "${PLAYBOOKS_DIR}" ]]; then
  echo "error: argument 1 (playbooks directory) is required"
  exit 1
fi

ansible-playbook "${PLAYBOOKS_DIR}/bashrc/undo.yml"
ansible-playbook "${PLAYBOOKS_DIR}/git/undo.credentials.yml"
ansible-playbook "${PLAYBOOKS_DIR}/git/undo.config.yml"
ansible-playbook "${PLAYBOOKS_DIR}/docker/undo.group.yml"
ansible-playbook "${PLAYBOOKS_DIR}/docker/undo.yml"
ansible-playbook "${PLAYBOOKS_DIR}/apt/autoremove.yml"
