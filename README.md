# Welcome to use_nix

This project was started to allow loading nix shell without leaving your current shell.

## Usage

    go get github.com/atrzaska/use_nix
    vim shell.nix # setup your nix shell in current working directory
    eval $(use_nix)
