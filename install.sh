#!/usr/bin/env bash

set -eo pipefail

COMMIT_HASH="$(curl -s "https://api.github.com/repos/chris-de-leon/infra-ubuntu/commits/master" | jq -erc '.sha')"

UBCTL_DIR="$(nix build --no-link "git+https://github.com/chris-de-leon/infra-ubuntu?rev=$COMMIT_HASH"#ubctl --print-out-paths)"

sudo ln -sfn "$UBCTL_DIR/bin/ubctl" "/usr/local/bin"
