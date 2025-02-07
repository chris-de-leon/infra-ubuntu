{ system ? builtins.currentSystem
  , pkgs ? import <nixpkgs> { inherit system; }
  ,
}:

pkgs.stdenv.mkDerivation rec {
  name = "ubctl";
  src = ../../.;
  propagatedBuildInputs = [
    pkgs.python311Packages.ansible-core
    pkgs.jq
  ];
  buildPhase = ''
    mkdir -p $out/bin

    cp -r $src/playbooks $out
    cp -r $src/nix/ubctl $out
    cp -r $src/scripts $out
    cp $src/ansible.cfg $out

    cp $src/nix/ubctl/main.sh $out/bin/${name}
    chmod +x $out/bin/${name}
  ''
};
