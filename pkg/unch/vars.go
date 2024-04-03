package unch

import (
	"errors"

	"encoding/base64"
)

/*
  - ErrIdenticalStr is returned when the input value 's' is equivalent to
    'message' during encoding.
  - ErrNotEncodedStr is returned when the output matches the input value 's'
    during decoding.
*/
var (
	ErrNotEncodedStr = errors.New("strings are not encoded")
	ErrIdenticalStr  = errors.New("the input strings are identical")

	b64 = base64.StdEncoding
)
