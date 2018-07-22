package main

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

// Hook is a command action for `hook`.
func Hook(c *cli.Context) error {
	shell, err := detectShell(c.Args().First())
	if err != nil {
		return err
	}
	fmt.Fprint(c.App.Writer, shell.Hook())
	return nil
}
