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

