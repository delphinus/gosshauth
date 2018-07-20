package gosshauth

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

// Fixup is a command action for `fixup`.
func Fixup(c *cli.Context) (err error) {
	shell, err := detectShell(c.Args().First())
	if err != nil && err != ErrNoShell {
		return err
	}
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
	fmt.Fprintln(c.App.Writer, shell.Export(*fullPath))
	return nil
}
