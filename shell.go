package gosshauth

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	// ErrUnknownShell means it is not a known shell.
	ErrUnknownShell = errors.New("unknown shell")
	// ErrNoShell means it lacks a shell string.
	ErrNoShell = errors.New("shell name needed: [zsh, bash]")
)

// Shell is an interface to export envs.
type Shell interface {
	Export(p string) string
	Hook() string
}

func detectShell(sh string) (shell Shell, err error) {
	switch sh {
	case "zsh":
		shell = ZSH
	case "bash":
		shell = BASH
	case "":
		err = ErrNoShell
	default:
		err = ErrUnknownShell
	}
	return
}

// me returns the full path for the executable.
func me() string {
	p, _ := exec.LookPath(os.Args[0])
	p, _ = filepath.Abs(p)
	return p
}
