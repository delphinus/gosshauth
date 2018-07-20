package gosshauth

import (
	"errors"
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

// Shell is an interface to export envs.
type Shell interface {
	Export(p string) string
}

// Fixup is a command action for `fixup`.
func Fixup(c *cli.Context) error {
	socks, err := SearchSockLinks()
	if err != nil {
		return err
	}
	err = SSHAuthSockEnv.FixWith(&socks.Newest().Path)
	if err == ErrLinkIsValid {
		return nil
	} else if err != nil {
		return err
	}
	fullPath, err := SSHAuthSockPath.FullPath()
	if err != nil {
		return err
	}
	if sh := c.Args().First(); sh != "" {
		shell, err := detectShell(sh)
		if err != nil {
			return err
		}
		fmt.Fprintln(c.App.Writer, shell.Export(*fullPath))
	}
	return nil
}

func detectShell(sh string) (shell Shell, err error) {
	switch sh {
	case "zsh":
		shell = ZSH
	case "bash":
		shell = BASH
	default:
		err = errors.New("unknown shell")
	}
	return
}
