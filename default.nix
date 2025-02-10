{ system ? builtins.currentSystem
, pkgs ? import <nixpkgs> { inherit system; }
,
}:

pkgs.buildGoModule rec {
  name = "src";
  src = ./.;
  vendorHash = "sha256-7SoKHH+tDJKhUQDoVwAzVZXoPuKNJEHDEyQ77BPEDQ0=";
}
