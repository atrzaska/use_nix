# Welcome to use_nix

This project was started to allow loading nix shell without leaving your current shell. Inspired by direnv https://github.com/direnv/direnv.

Under the hood it will run `nix-shell --show-trace --run "env"`. It will compute a diff between your current env state and output nix env. Finally it will output a script that is ready for evaluting in your current shell.

## Usage

    go get github.com/atrzaska/use_nix
    vim shell.nix # setup your nix shell in current working directory
    eval $(use_nix)
