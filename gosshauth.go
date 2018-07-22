package main

import (
	"os"
)

func main() {
	if err := NewApp().Run(os.Args); err != nil {
		// It never reaches here because urfave/cli handles it.
		os.Exit(1)
	}
}
