{
  description = "ft_linear_regression — linear regression with gradient descent";

  inputs = {
    nixpkgs.url = "https://channels.nixos.org/nixpkgs-unstable/nixexprs.tar.zst";
  };

  outputs = inputs: {
    devShells = builtins.mapAttrs (system: pkgs: {
      default = pkgs.mkShell {
        packages = [
          pkgs.go
          pkgs.gnumake

          # neovim go tooling
          pkgs.gopls # lsp
          pkgs.golangci-lint # linter
          pkgs.gofumpt # formatter
          pkgs.gotools # goimports
          pkgs.delve # debugger
        ];
      };
    }) inputs.nixpkgs.legacyPackages;
  };
}
