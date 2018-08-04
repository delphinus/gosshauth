package main

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
	if tmpl := c.String("format"); tmpl != "" {
		txt, err := socks.WithTemplate(tmpl)
		if err != nil {
			return fmt.Errorf("error in the supplied template: %s", err)
		}
		fmt.Fprintf(c.App.Writer, "%s", txt)
	} else {
		fmt.Fprintf(c.App.Writer, "%s", socks)
	}
	return nil
}
