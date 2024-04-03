// Package unch provides functions for encoding and decoding messages using
// invisible Unicode characters.
package unch

import "strings"

// Encode the given message and embeds it into the provided string. It uses
// invisible Unicode characters for masking.
func Encode(message, s string) (string, error) {
	// Compare input & output strings
	if message == s {
		return s, ErrIdenticalStr
	}

	// Encode the message to base64
	message = b64.EncodeToString([]byte(message))

	var encoded strings.Builder
	for _, char := range message {
		// Add invisible Unicode char to encode
		encoded.WriteRune(char + E0100)
	}

	return s + encoded.String(), nil
}

// Decode the encoded message from the provided string. It extracts the hidden
// message and returns it.
func Decode(s string) (string, error) {
	var decoded strings.Builder
	for _, char := range s {
		// Remove the invisible Unicode chars to decode the msg
		decoded.WriteRune(char - E0100)
	}

	// Remove non-ASCII chars
	stripped := strings.Map(isRuneASCII, decoded.String())

	// Decode the msg from base64
	message, err := b64.DecodeString(stripped)
	if err != nil {
		return s, err
	}

	output := string(message)

	// Compare input & output strings
	if s == output || output == "" {
		return s, ErrNotEncodedStr
	}

	return output, nil
}
