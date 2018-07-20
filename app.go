package gosshauth

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

// NewApp makes CLI app.
func NewApp() *cli.App {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "%v version v%v (%s)\n",
			c.App.Name, c.App.Version, GitCommit)
	}
	return &cli.App{
		Version: Version,
		Usage:   "Detect $SSH_AUTH_SOCK and fix the symlink",
		Authors: []*cli.Author{
			{Name: "JINNOUCHI Yasushi", Email: "delphinus@remora.cx"},
		},
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List up existent & accessible sock files",
				Action:  List,
			},
			{
				Name:    "fixup",
				Aliases: []string{"f"},
				Usage:   "Fix up the link for sock files and export $SSH_AUTH_SOCK",
				Action:  Fixup,
			},
		},
	}
}
