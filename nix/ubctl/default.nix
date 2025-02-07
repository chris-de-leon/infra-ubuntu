{ system ? builtins.currentSystem
  , pkgs ? import <nixpkgs> { inherit system; }
  ,
}:

pkgs.stdenv.mkDerivation rec {
  pname = "ubctl";
  src = ../../.;

  entrypoint = pkgs.writeShellApplication {
    name = "main";
    runtimeInputs = [ pkgs.python311Packages.ansible-core pkgs.jq ];
    text = builtins.readFile ./main.sh;
  };

  buildPhase = ''
    mkdir -p $out/bin

    cp -r $src/nix/ubctl/cmd $out
    cp -r $src/playbooks $out
    cp $src/ansible.cfg $out

    cp ${entrypoint}/bin/main $out/bin/${name}
    chmod +x $out/bin/${name}
    chmod -R +x $out/cmd
  '';
}
