package unch

func isRuneASCII(r rune) rune {
	if r <= 127 {
		return r
	}

	// non-ASCII char
	return -1
}
