#!/usr/bin/env bash

set -eo pipefail

SCRIPT_DIR="$(cd -- "$(dirname "$0")" >/dev/null 2>&1 && pwd -P)"
ASSETS_DIR="$(realpath "$SCRIPT_DIR/..")"

op1="${1:-}"
op2="${2:-}"

case "$op1" in
welcome)
  echo "Welcome! If you are seeing this message, then you have successfully installed ubctl!"
  echo -e "
'||'  '|' '||''|.     ..|'''.| |''||''| '||'      
 ||    |   ||   ||  .|'     '     ||     ||       
 ||    |   ||'''|.  ||            ||     ||       
 ||    |   ||    || '|.      .    ||     ||       
  '|..'   .||...|'   ''|....'    .||.   .||.....| 
"
  ;;
shell)
  cd "$ASSETS_DIR" && nix develop --show-trace --no-write-lock-file "${@:2}"
  ;;
vm)
  case "$op2" in
  init)
    "$ASSETS_DIR/cmd/vm/init.sh" "$ASSETS_DIR/playbooks" "${@:3}"
    ;;
  reset)
    "$ASSETS_DIR/cmd/vm/undo.sh" "$ASSETS_DIR/playbooks" "${@:3}"
    ;;
  *)
    echo "Invalid option: $op1 $op2"
    echo "Usage: $0 $op1 {init|reset}"
    exit 1
    ;;
  esac
  ;;
dotfiles)
  case "$op2" in
  pull)
    "$ASSETS_DIR/cmd/dotfiles/init.sh" "${@:3}"
    ;;
  purge)
    "$ASSETS_DIR/cmd/dotfiles/undo.sh" "${@:3}"
    ;;
  *)
    echo "Invalid option: $op1 $op2"
    echo "Usage: $0 $op1 {pull|purge}"
    exit 1
    ;;
  esac
  ;;
*)
  echo "Invalid option: $op1"
  echo "Usage: $0 {shell|vm|dotfiles}"
  exit 1
  ;;
esac
