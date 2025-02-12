{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.zsh
    pkgs.git
    pkgs.openssl
  ];

  shellHook = ''
    alias vim="nvim"
    alias vi="nvim"
    alias v="nvim"

    # Modify this
    export PS1="\[\e[35m\]\w \$ \[\e[0m\]"

    echo ""
    echo "The default environment is ready!"
    echo ""
  '';
}
