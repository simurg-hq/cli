// Package cli implements the commands of the simurg command line tool.
package cli

import (
	"fmt"
	"io"
)

// command is the name of the binary, used in usage and version output.
const command = "simurg"

// Exit codes.
const (
	exitOK    = 0
	exitError = 1
	exitUsage = 2
)

// Run executes the command in args (without the program name) and returns
// the process exit code. version is the build version of the binary.
func Run(args []string, version string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		usage(stderr)
		return exitUsage
	}

	switch args[0] {
	case "version":
		if len(args) > 1 {
			printErr(stderr, "version takes no arguments\n")
			return exitUsage
		}
		if _, err := fmt.Fprintf(stdout, "%s %s\n", command, version); err != nil {
			printErr(stderr, "write version: %v\n", err)
			return exitError
		}
		return exitOK
	default:
		printErr(stderr, "unknown command %q\n", args[0])
		usage(stderr)
		return exitUsage
	}
}

func usage(stderr io.Writer) {
	printErr(stderr, "usage: %s <command>\n\ncommands:\n  version  print the version\n", command)
}

// printErr writes to stderr. If stderr itself fails there is nowhere left to
// report it, so the error is dropped.
func printErr(stderr io.Writer, format string, args ...any) {
	fmt.Fprintf(stderr, format, args...)
}
