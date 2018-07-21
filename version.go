package gosshauth

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "%v version v%v (%s)\n",
			c.App.Name, c.App.Version, GitCommit)
	}
}

// Version is needed for selfupdate.
const Version = "0.0.4"

// GitCommit is needed for selfupdate.  This will be set in Makefile task.
var GitCommit = ""
