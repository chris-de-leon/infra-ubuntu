# Windows Setup

## Docker Desktop

Follow the official installation instructions [here](https://www.docker.com/products/docker-desktop/) to install Docker Desktop if you haven't already.

## WSL

Follow the official installation instructions [here](https://learn.microsoft.com/en-us/windows/wsl/install) to install WSL 2 if you haven't already.

### Starting Fresh

1. Open a command prompt
1. List the distros available for download: `wsl -l -o`
1. Install the distro of your choice: `wsl --install -d <distro>`
1. Enter a username and password for the distro

### Upgrading

1. Open a command prompt
1. Make a backup of your files
1. List the distros on your machine: `wsl -l`
1. Uninstall any unnecessary distros: `wsl --unregister <distro>`
1. List the distros available for download: `wsl -l -o`
1. Install the distro of your choice: `wsl --install -d <distro>`
1. Enter a username and password for the distro


