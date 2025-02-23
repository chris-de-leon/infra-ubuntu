# https://github.com/NixOS/nixpkgs/commits/master
{
  inputs = {
    nixpkgs.url = "https://github.com/NixOS/nixpkgs/archive/4bbb73beb26f5152a87d3460e6bf76227841ae21.tar.gz";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      rec {
        formatter = pkgs.nixpkgs-fmt;

        # Configures tmux to use interactive bash:
        #  - https://unix.stackexchange.com/a/663023
        #  - https://askubuntu.com/a/746846
        #
        wrappedTmux = pkgs.symlinkJoin {
          name = "tmux";
          paths = [ pkgs.tmux ];
          buildInputs = [ pkgs.makeWrapper ];
          postBuild = ''
            wrapProgram $out/bin/tmux \
              --set "SHELL" "${pkgs.bashInteractive}/bin/bash" \
              --add-flags "-u"
          '';
        };

        devShells = {
          # NOTE: we need to add the starship eval command in two places. The first place
          # is the ~/.bashrc file as per the docs, and the second place is the shell hook
          # below since ~/.bashrc won't be sourced when we run `nix develop`.
          default = pkgs.mkShell rec {
            packages = [
              pkgs.bashInteractive
              pkgs.shellcheck
              pkgs.unzipNLS
              pkgs.starship
              pkgs.ripgrep
              pkgs.lazygit
              pkgs.neovim
              wrappedTmux
              pkgs.fzf
              pkgs.vim
              pkgs.fd
              pkgs.jq
            ];
            shellHook = ''
              eval "$(starship init bash)"
            '';
          };
        };
      }
    );
}

