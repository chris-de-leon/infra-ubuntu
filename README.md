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

1. Install the `ubctl` CLI:

    ```sh
    curl -sfL https://raw.githubusercontent.com/chris-de-leon/infra-ubuntu/refs/heads/master/install.sh | bash
    ```

1. If you'd like to uninstall the `ubctl` CLI later, then you can run:

    ```sh
    sudo rm /usr/bin/local/ubctl && rm -rf ~/.config/ubctl
    ```

## Usage

1. Setup a new Ubuntu VM with Docker, dev tools (e.g. Starship), and Git:

    ```sh
    # After running this, you should exit and re-enter the VM
    ubctl vm init
    ```

1. Pull the latest dotfiles:

    ```sh
    ubctl dotfiles pull
    ```

1. Enter a fully-configured dev shell:

    ```sh
    ubctl shell
    ```

## Upgrading


1. Upgrade the CLI to the latest version:

    ```sh
    curl -sfL https://raw.githubusercontent.com/chris-de-leon/infra-ubuntu/refs/heads/master/install.sh | bash
    ```

1. Pull the latest dotfiles:

    ```sh
    ubctl dotfiles pull
    ```
