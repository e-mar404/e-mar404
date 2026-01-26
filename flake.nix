{
  description = "All things needed for development on this site";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-25.05";
  };
  
  outputs = { self, nixpkgs }: 

  let 
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
  in
  {
    devShells.${system}.default = pkgs.mkShell {
      buildInputs = with pkgs; [
        go
        hugo
        gcc
      ];
    };
  };
}
