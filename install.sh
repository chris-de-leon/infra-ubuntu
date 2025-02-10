#!/usr/bin/env bash

set -eo pipefail

# TODO: download binary from repo's releases page, add it to /usr/local/bin, chmod +x, and print version

# COMMIT_HASH="$(curl -s "https://api.github.com/repos/chris-de-leon/infra-ubuntu/commits/master" | jq -erc '.sha')"
#
# UBCTL_DIR="$(nix build --no-link "git+https://github.com/chris-de-leon/infra-ubuntu?rev=${COMMIT_HASH}"#ubctl --print-out-paths)"
#
# cat <<EOF | sudo tee /usr/local/bin/ubctl >/dev/null
# #!/usr/bin/env bash
# exec "${UBCTL_DIR}/bin/ubctl" "\$@"
# EOF
#
# sudo chmod +x /usr/local/bin/ubctl
#
# ubctl welcome

echo "This script is currently under development"
