package main

import (
	"bufio"
	"fmt"
	"os"
)

func showBanner() {
	version := AppVersion
	if BuildCommit != "" {
		version += "-" + BuildCommit
	}

	fmt.Fprintf(os.Stderr, banner, version)
}

func showHelps() {
	showBanner()
	fmt.Fprintf(os.Stderr, "\n\nUsage: %s", usage)
	fmt.Fprintf(os.Stderr, "\n\nOptions: %s", options)
	fmt.Fprintf(os.Stderr, "\n\nExamples: %s\n", examples)
}

func isStdin() bool {
	stdin, _ := os.Stdin.Stat()

	return (stdin.Mode() & os.ModeCharDevice) == 0
}

func getStdin() string {
	stdin := bufio.NewScanner(os.Stdin)
	if isStdin() && stdin.Scan() {
		return stdin.Text()
	}

	return ""
}
