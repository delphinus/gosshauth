package gosshauth

type zsh struct{}

// ZSH is a shell `zsh`.
var ZSH = zsh{}

func (zsh) Export(p string) string {
	return "export SSH_AUTH_SOCK=" + p
}
