package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	ErrNoOp        = errors.New("no-op")
	ErrNoText      = errors.New("no PLAINTEXT provided")
	ErrTooManyArgs = errors.New("too many arguments")
)

func showErr(err error) {
	fmt.Printf("error: %s\n", err.Error())

	flag.Usage()
	os.Exit(1)
}
