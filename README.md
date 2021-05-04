# Welcome to use_nix

This project was started to allow loading nix shell without leaving your current shell. Inspired by direnv https://github.com/direnv/direnv.

Under the hood it will run `nix-shell --show-trace --run "env"`. It will compute a diff between your current env state and output nix env. Finally it will output a script that is ready for evaluting in your current shell.

## Installation

    go get -u github.com/atrzaska/use_nix

## Usage

To load nix shell of your project make sure you have properly setup `nix` on your system and `shell.nix` on your project. After that you can run:

    eval $(use_nix)

This command will evaluate your `shell.nix` and setup your enviornment variables in your current shell.
