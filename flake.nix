{
  description = "All things needed for development on this site";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-25.05";
  };
  
  outputs = { self, nixpkgs }: 

    let 
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
   in
    {
      devShells = nixpkgs.lib.genAttrs systems (system: 
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in {
          default = pkgs.mkShell {
            packages = with pkgs; [ go hugo stdenv.cc ];
          };
        }
      );
    };
}
