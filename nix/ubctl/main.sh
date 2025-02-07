#!/usr/bin/env bash

set -eo pipefail

SCRIPT_DIR="$(cd -- "$(dirname "$0")" >/dev/null 2>&1 && pwd -P)"
ASSETS_DIR="$(realpath "$SCRIPT_DIR/..")"

INFRA_UBUNTU_HASH="${INFRA_UBUNTU_REV:-}"
INFRA_UBUNTU_OWNR="chris-de-leon"
INFRA_UBUNTU_REPO="infra-ubuntu"
get_rev() {
  if [ -z "$INFRA_UBUNTU_HASH" ]; then
    curl -s "https://api.github.com/repos/$INFRA_UBUNTU_OWNR/$INFRA_UBUNTU_REPO/commits/master" | jq -erc '.sha'
  else
    echo "$INFRA_UBUNTU_HASH"
  fi
}

op1="$1"
case "$op1" in
upgrade)
  "$ASSETS_DIR/cmd/upgrade.sh" "$ASSETS_DIR/playbooks"
  ;;
shell)
  "$ASSETS_DIR/cmd/shell.sh" "https://github.com/$INFRA_UBUNTU_OWNR/$INFRA_UBUNTU_REPO?rev=$(get_rev)"
  ;;
rev)
  get_rev
  ;;
dotfiles)
  op2="$2"
  case "$op2" in
  upgrade)
    "$ASSETS_DIR/cmd/dotfiles/init.sh" -f
    ;;
  remove)
    "$ASSETS_DIR/cmd/dotfiles/undo.sh" -f
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
  echo "Usage: $0 {upgrade|shell|rev|dotfiles}"
  exit 1
  ;;
esac
