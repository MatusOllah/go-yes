package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

const ver string = "1.0.0"

func main() {
	help := flag.Bool("help", false, "display this help and exit")
	version := flag.Bool("version", false, "output version information and exit")
	flag.Parse()

	if *help {
		fmt.Fprintln(os.Stderr, "go-yes repeatedly outputs a line with specified string, or 'y'.")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *version {
		fmt.Fprintf(os.Stderr, "go-yes version %s\n", ver)
		fmt.Fprintf(os.Stderr, "Go version %s\n", runtime.Version())
		os.Exit(0)
	}

	for {
		if len(os.Args) >= 2 {
			fmt.Fprintln(os.Stdout, strings.Join(os.Args[1:], " "))
		} else {
			fmt.Fprintln(os.Stdout, "y")
		}
	}
}
