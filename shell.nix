{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/tarball/21.11") { }
}:
pkgs.mkShell {
  buildInputs = [
    # GO
    pkgs.go

    # RUST
    pkgs.rustc
    pkgs.cargo
  ];
}
