package main

import (
	"flag"
)

func init() {
	opts = new(Options)

	flag.BoolVar(&opts.Decode, "d", false, "")
	flag.BoolVar(&opts.Lorem, "l", false, "")
	flag.BoolVar(&opts.Lorem, "lorem", false, "")
	flag.BoolVar(&opts.NoNewline, "n", false, "")

	flag.Usage = func() { showHelps() }
	flag.Parse()
}
