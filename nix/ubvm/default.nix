{ system ? builtins.currentSystem
  , pkgs ? import <nixpkgs> { inherit system; }
  ,
}:

pkgs.stdenv.mkDerivation rec {
  name = "ubvm";
  src = ../../.;
  buildInputs = [
    pkgs.python311Packages.ansible-core
    pkgs.jq
  ];
  buildPhase = ''
    mkdir -p $out/bin

    cp -r $src/playbooks $out
    cp -r $src/nix/ubvm $out
    cp $src/ansible.cfg $out

    cp $src/nix/ubvm/main.sh $out/bin/${name}
    chmod +x $out/bin/${name}
  '';
}
