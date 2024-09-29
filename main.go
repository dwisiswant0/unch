package main

import (
	"flag"
	"fmt"
)

func main() {
	var args []string = flag.Args()
	var err error
	var output string

	switch {
	case len(args) > 0 || isStdin():
		stdin := getStdin()
		if stdin != "" {
			args = append([]string{stdin}, args...)
		}

		if len(args) > 2 {
			showErr(ErrTooManyArgs)
		}

		switch {
		case len(args) <= 0:
			showErr(ErrNoOp)
		case opts.Decode:
			output, err = opts.decode(args[0])
		case opts.Lorem:
			output, err = opts.encode(args...)
		case len(args) > 1:
			if args[1] == "" && !opts.Lorem {
				showErr(ErrTooManyArgs)
			}

			output, err = opts.encode(args...)
		}
	default:
		showErr(ErrNoOp)
	}

	if err != nil {
		showErr(err)
	}

	if !opts.NoNewline {
		output += "\n"
	}

	fmt.Print(output)
}
