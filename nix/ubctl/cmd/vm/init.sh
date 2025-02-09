#!/usr/bin/env bash

set -eo pipefail

PLAYBOOKS_DIR="${1:-}"
if [[ -z "${PLAYBOOKS_DIR}" ]]; then
  echo "error: argument 1 (playbooks directory) is required"
  exit 1
fi

GH_TOKEN="${GH_TOKEN:-}"
GH_UNAME="${GH_UNAME:-}"
GH_EMAIL="${GH_EMAIL:-}"
GH_NAME="${GH_NAME:-}"

if [[ -z "${GH_UNAME}" ]]; then read -rp "Enter your GitHub username: " GH_UNAME; fi
if [[ -z "${GH_TOKEN}" ]]; then read -rp "Enter your GitHub token: " GH_TOKEN; fi
if [[ -z "${GH_EMAIL}" ]]; then read -rp "Enter your GitHub email: " GH_EMAIL; fi
if [[ -z "${GH_NAME}" ]]; then read -rp "Enter your GitHub name: " GH_NAME; fi

ansible-playbook "${PLAYBOOKS_DIR}/apt/update-upgrade.yml"
ansible-playbook "${PLAYBOOKS_DIR}/docker/init.yml"
ansible-playbook "${PLAYBOOKS_DIR}/docker/init.group.yml"
ansible-playbook "${PLAYBOOKS_DIR}/git/init.config.yml" -e "email=${GH_EMAIL} name=${GH_NAME}"
ansible-playbook "${PLAYBOOKS_DIR}/git/init.credentials.yml" -e "token=${GH_TOKEN} uname=${GH_UNAME}"
ansible-playbook "${PLAYBOOKS_DIR}/bashrc/init.yml"
