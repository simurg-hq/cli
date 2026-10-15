// Command simurg is the command line tool of the ping service.
package main

import (
	"os"

	"github.com/simurg-hq/cli/internal/cli"
)

// version is set at build time:
//
//	go build -ldflags "-X main.version=v0.1.0"
var version = "dev"

func main() {
	os.Exit(cli.Run(os.Args[1:], version, os.Stdout, os.Stderr))
}
