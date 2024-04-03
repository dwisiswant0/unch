package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/dwisiswant0/unch/pkg/unch"
)

func (opts *Options) encode(s ...string) (string, error) {
	var message string = s[0]
	var plaintext string

	if opts.Lorem {
		plaintext = gofakeit.LoremIpsumSentence(5)
	} else if len(s) > 1 {
		plaintext = s[1]
	}

	return unch.Encode(message, plaintext)
}

func (opts *Options) decode(s string) (string, error) {
	return unch.Decode(s)
}
