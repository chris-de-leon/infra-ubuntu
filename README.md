# Infra-Ubuntu

## Archive Notice

This repo has been archived in favor of [dotfiles](https://github.com/chris-de-leon/dotfiles)

## Setup

Before starting, you'll need to create a Github [fine-grained personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token).

### Windows / WSL / Ubuntu

1. Follow the WSL setup guide [here](./docs/wsl.md)

1. Install Alacritty and configure it using the guide [here](./docs/alacritty.windows.md)

1. Open an Alacritty terminal, and you'll be placed in a WSL Ubuntu shell

### Ubuntu Desktop

1. Install Alacritty from the App Center and configure it using the guide [here](./docs/alacritty.ubuntu.md)

1. Install Multipass from the App Center and launch the Multipass app

1. Open an Alacritty terminal and create a new multipass VM:

   ```sh
   # Values are for display purposes only
   multipass launch 24.04 --name=dev --cpus=12 --memory=30G --disk=100G
   ```

1. Enter the VM:

   ```sh
   multipass shell dev
   ```

## Installing

Once you're inside the VM, follow the steps below:

1. Install Nix:

   ```sh
   curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install --no-confirm
   ```

1. Install the `ubctl` CLI:

   ```sh
   curl -sfL https://raw.githubusercontent.com/chris-de-leon/infra-ubuntu/refs/heads/master/install.sh | bash
   ```

1. Initialize the VM using the command below:

   ```sh
   # After running this, you should exit and re-enter the VM
   ubctl vm init --gh-username "your-github-username" --gh-token "your-github-token" --gh-email "your.email@mail.com" --gh-name "your-name"
   ```

1. Once the VM is initialized, you can enter a fully-configured dev shell:

   ```sh
   ubctl shell
   ```

## Upgrading

1. Clean up the VM:

   ```sh
   ubctl vm undo && ubctl clean && nix-collect-garbage
   ```

1. Upgrade the CLI to the latest version:

   ```sh
   curl -sfL https://raw.githubusercontent.com/chris-de-leon/infra-ubuntu/refs/heads/master/install.sh | bash
   ```

1. Setup the VM:

   ```sh
   # After running this, you should exit and re-enter the VM
   ubctl vm init --gh-username "your-github-username" --gh-token "your-github-token" --gh-email "your.email@mail.com" --gh-name "your-name"
   ```

## Uninstalling

1. Revert the VM back to its original state:

   ```sh
   ubctl vm undo
   ```

1. Clean up all artifacts created by the CLI:

   ```sh
   ubctl clean
   ```

1. Remove the executable:

   ```sh
   sudo rm /usr/local/bin/ubctl
   ```

1. Uninstall Nix:

   ```sh
   /nix/nix-installer uninstall --no-confirm
   ```
