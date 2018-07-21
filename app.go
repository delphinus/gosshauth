package gosshauth

import (
	"time"

	"gopkg.in/urfave/cli.v2"
)

// NewApp makes CLI app.
func NewApp() *cli.App {
	return &cli.App{
		Version: Version,
		Usage:   "Detect $SSH_AUTH_SOCK and fix the symlink",
		Authors: []*cli.Author{
			{Name: "JINNOUCHI Yasushi", Email: "delphinus@remora.cx"},
		},
		Compiled: time.Unix(compileTime, 0),
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List up existent & accessible sock files",
				Action:  actionFunc(List),
			},
			{
				Name:      "fixup",
				Aliases:   []string{"f"},
				Usage:     "Fix up the link for sock files",
				ArgsUsage: "[(none), zsh, bash]",
				Description: "Check $SSH_AUTH_SOCK and validate it.  " +
					"When you supply a shell name, print out export setting for it " +
					"(only if needed).",
				Action: actionFunc(Fixup),
			},
			{
				Name:      "hook",
				Usage:     "Show hook script for shells",
				ArgsUsage: "[zsh, bash]",
				Description: "Show hook script for supplied shells.  " +
					"Use as `eval $(gosshauth hook zsh)` to set hooks.",
				Action: actionFunc(Hook),
			},
			{
				Name:   "selfupdate",
				Usage:  "Update the binary itself",
				Action: actionFunc(Selfupdate),
			},
		},
	}
}

func actionFunc(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		if err := f(c); err != nil {
			if _, ok := err.(cli.ExitCoder); ok {
				return err
			}
			return cli.Exit(err, 1)
		}
		return nil
	}
}
