package main

type zsh struct{}

// ZSH is a shell `zsh`.
var ZSH = zsh{}

func (zsh) Export(p string) string {
	return "export SSH_AUTH_SOCK=" + p
}

func (zsh) Hook() string {
	return `_sshauthsock_hook() {
    eval "$("` + Me() + `" fixup zsh)";
}
typeset -ag precmd_functions;
if [[ -z ${precmd_functions[(r)_sshauthsock_hook]} ]]; then
    precmd_functions+=_sshauthsock_hook;
fi
`
}
