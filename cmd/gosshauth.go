package main

import (
	"os"

	"github.com/delphinus/gosshauth"
)

func main() {
	if err := gosshauth.NewApp().Run(os.Args); err != nil {
		// It never reaches here because urfave/cli handles it.
		os.Exit(1)
	}
}
