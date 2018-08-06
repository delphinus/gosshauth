package main

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/urfave/cli.v2"
)

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		rev := GitCommit
		if rev == "" {
			rev = "dev"
		}
		compiled := ""
		if c.App.Compiled.Unix() == 0 {
			compiled = "development"
		} else {
			compiled = c.App.Compiled.Format(time.RFC3339)
		}
		fmt.Fprintf(c.App.Writer, "%v version v%v (%s) build time: %s\n",
			c.App.Name, c.App.Version, rev, compiled)
	}
}

// Version is needed for selfupdate.
const Version = "0.0.13"

// GitCommit is needed for selfupdate.  This will be set in Makefile task.
var GitCommit = ""

// CompileTime is unixtime when this binary is compiled.  This will be set
// in Makefile task.
var CompileTime = ""
var compileTime int64

func init() { compileTime, _ = strconv.ParseInt(CompileTime, 10, 64) }
