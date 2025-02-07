# Infra-Ubuntu (Ubuntu Desktop v24.04 LTS, noble)

## Setup

1. Create a Github [fine-grained personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token)

1. Use the App Center to install Alacritty and set it up using the guide [here](./docs/alacritty.ubuntu.md)

1. Use the App Center to install Multipass and launch the Multipass app

1. Open an Alacritty terminal and create a new multipass VM:

    ```sh
    multipass launch 24.04 --name=dev --cpus=12 --memory=30G --disk=100G
    ```

1. Enter the VM:

    ```sh
    multipass shell dev
    ```

1. Install Nix:

    ```sh
    curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install --no-confirm
    ```

1. Get the latest commit hash of this repo:

    ```sh
    INFRA_UBUNTU_REV="$(curl -s "https://api.github.com/repos/chris-de-leon/infra-ubuntu/commits/master" | grep -m 1 '"sha":' | awk -F '"' '{print $4}')" && echo "$INFRA_UBUNTU_REV"
    ```

1. Install the `ubctl` CLI:

    ```sh
    nix run "git+https://github.com/chris-de-leon/infra-ubuntu?rev=$INFRA_UBUNTU_REV"#install
    ```

1. If you'd like to uninstall the CLI later, then you can run:

    ```sh
    nix run "git+https://github.com/chris-de-leon/infra-ubuntu?rev=$INFRA_UBUNTU_REV"#uninstall
    ```
1. Source `~/.bashrc`:

    ```sh
    . ~/.bashrc
    ```

## Usage

1. Setup a new Ubuntu VM with Docker, dev tools (e.g. Starship), and Git:

    ```sh
    ubctl vm init && ~/.bashrc
    ```

1. Pull the latest dotfiles:

    ```sh
    ubctl dotfiles upgrade
    ```

1. Enter a fully-configured dev shell:

    ```sh
    ubctl shell
    ```

## Upgrading


1. Upgrade the CLI to the latest version:

    ```sh
    ubctl upgrade && . ~/.bashrc
    ```

1. Pull the latest dotfiles:

    ```sh
    ubctl dotfiles upgrade
    ```
