package gosshauth

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

// Fixup is a command action for `fixup`.
func Fixup(c *cli.Context) error {
	if socks, err := SearchSockLinks(); err != nil {
		return err
	} else if err := FixSSHAuthSock(socks); err == ErrLinkIsValid {
		return nil
	} else if err != nil {
		return err
	}
	fullPath, err := sshAuthSockPath.FullPath()
	if err != nil {
		return err
	}
	fmt.Fprintln(c.App.Writer, "export SSH_AUTH_SOCK="+*fullPath)
	return nil
}
