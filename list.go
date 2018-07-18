package gosshauth

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

// List is a command action for `list`.
func List(c *cli.Context) error {
	socks, err := SearchSockLinks()
	if err != nil {
		return fmt.Errorf("error in SockLinks: %v", err)
	}
	fmt.Fprintf(c.App.Writer, "%s", socks)
	return nil
}
