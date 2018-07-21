package gosshauth

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/urfave/cli.v2"
)

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "%v version v%v (%s) build time: %s\n",
			c.App.Name, c.App.Version, GitCommit, c.App.Compiled.Format(time.RFC3339))
	}
}

// Version is needed for selfupdate.
const Version = "0.0.4"

// GitCommit is needed for selfupdate.  This will be set in Makefile task.
var GitCommit = ""

// CompileTime is unixtime when this binary is compiled.  This will be set
// in Makefile task.
var CompileTime = ""
var compileTime int64

func init() { compileTime, _ = strconv.ParseInt(CompileTime, 10, 64) }
