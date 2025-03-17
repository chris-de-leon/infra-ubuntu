# https://github.com/NixOS/nixpkgs/commits/master
{
  inputs = {
    nixpkgs.url = "https://github.com/NixOS/nixpkgs/archive/5996243e793c19a0933c8df5869b79088279343d.tar.gz";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      rec {
        formatter = pkgs.nixpkgs-fmt;

        devShells = {
          dev = pkgs.mkShell rec {
            packages = [
              pkgs.goreleaser
              pkgs.ansible
              pkgs.nodejs
              pkgs.go
              pkgs.gh
            ];
          };
        };

        packages = {
          ubctl = pkgs.callPackage ./default.nix {
            system = system;
            pkgs = pkgs;
          };
        };
      }
    );
}

