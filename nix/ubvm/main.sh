#!/usr/bin/env bash

set -eo pipefail

SCRIPT_DIR="$(cd -- "$(dirname "$0")" >/dev/null 2>&1 && pwd -P)"
ASSETS_DIR="$(realpath "$SCRIPT_DIR/..")"

op="$1"
case "$op" in
init)
  GH_TOKEN="${GH_TOKEN:-}"
  GH_UNAME="${GH_UNAME:-}"
  GH_EMAIL="${GH_EMAIL:-}"
  GH_NAME="${GH_NAME:-}"

  if [ -z "$GH_UNAME" ]; then read -rp "Enter your GitHub username: " GH_UNAME; fi
  if [ -z "$GH_TOKEN" ]; then read -rp "Enter your GitHub token: " GH_TOKEN; fi
  if [ -z "$GH_EMAIL" ]; then read -rp "Enter your GitHub email: " GH_EMAIL; fi
  if [ -z "$GH_NAME" ]; then read -rp "Enter your GitHub name: " GH_NAME; fi

  ansible-playbook "$ASSETS_DIR/playbooks/apt/update-upgrade.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/docker/init.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/docker/init.group.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/git/init.config.yml" -e "email=$GH_EMAIL name=$GH_NAME"
  ansible-playbook "$ASSETS_DIR/playbooks/git/init.credentials.yml" -e "token=$GH_TOKEN uname=$GH_UNAME"
  ansible-playbook "$ASSETS_DIR/playbooks/bashrc/init.version.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/bashrc/init.yml"

  . "$HOME/.bashrc"
  ;;
undo)
  ansible-playbook "$ASSETS_DIR/playbooks/bashrc/undo.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/bashrc/undo.version.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/git/undo.credentials.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/git/undo.config.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/docker/undo.group.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/docker/undo.yml"
  ansible-playbook "$ASSETS_DIR/playbooks/apt/autoremove.yml"

  . "$HOME/.bashrc"
  ;;
*)
  echo "Invalid option: $op"
  echo "Usage: $0 {init|undo}"
  exit 1
  ;;
esac
