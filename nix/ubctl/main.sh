#!/usr/bin/env bash

set -eo pipefail

SCRIPT_DIR="$(cd -- "$(dirname "$0")" >/dev/null 2>&1 && pwd -P)"
ASSETS_DIR="$(realpath "$SCRIPT_DIR/..")"

INFRA_UBUNTU_OWNR="chris-de-leon"
INFRA_UBUNTU_REPO="infra-ubuntu"
get_rev() {
  if [ -z "$INFRA_UBUNTU_REV" ]; then
    curl -s "https://api.github.com/repos/$INFRA_UBUNTU_OWNR/$INFRA_UBUNTU_REPO/commits/master" | jq -erc '.sha'
  else
    echo "$INFRA_UBUNTU_REV"
  fi
}

op1="$1"
op2="$2"
case "$op1" in
upgrade)
  ansible-playbook "$(realpath "$ASSETS_DIR/playbooks/bashrc/init.version.yml")" && . "$HOME/.bashrc"
  ;;
shell)
  nix develop "https://github.com/$INFRA_UBUNTU_OWNR/$INFRA_UBUNTU_REPO?rev=$(get_rev)"
  ;;
rev)
  get_rev
  ;;
dotfiles)
  case "$op2" in
  upgrade)
    bash "$(realpath "$ASSETS_DIR/scripts/dotfiles/init.sh")" -f
    ;;
  remove)
    bash "$(realpath "$ASSETS_DIR/scripts/dotfiles/undo.sh")" -f
    ;;
  *)
    echo "Invalid option: $op1 $op2"
    echo "Usage: $0 $op1 {upgrade|remove}"
    exit 1
    ;;
  esac
  ;;
*)
  echo "Invalid option: $op1"
  echo "Usage: $0 {upgrade|shell|dotfiles}"
  exit 1
  ;;
esac
