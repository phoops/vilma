let
nixpkgs = builtins.fetchTarball {
  name = "nixpkgs-for-vilma-2022-01-22";
  url = "https://github.com/nixos/nixpkgs/archive/fad04722fc3d692e3511e58e337ec9fa627f5ba5.tar.gz";
  # Hash obtained using `nix-prefetch-url --unpack <url>`
  sha256 = "sha256:04s1gsdnvzw57kidarb52grqrr6ayg0gmy9pifkalyvkdhkyrxdr";
};
in
{ pkgs ? import nixpkgs {}}:
pkgs.mkShell {
    buildInputs = [
      pkgs.protoc-gen-go
      pkgs.grpc-tools
    ];
}