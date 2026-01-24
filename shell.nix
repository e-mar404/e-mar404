{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    hugo
    gcc
    tailwindcss
    nodePackages.postcss
    autoprefixer
    nodejs_24
  ];
}
