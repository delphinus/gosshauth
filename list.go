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
	tmpl := c.String("format")
	if tmpl == "" {
		return socks.Print(c.App.Writer)
	}
	return socks.PrintWithTemplate(c.App.Writer, tmpl)
}
