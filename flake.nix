{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
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

