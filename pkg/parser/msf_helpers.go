package parser

import (
"strconv"
"unicode/utf8"
)

// Helper functions

// Unquote interprets s as a single-quoted, double-quoted,
// or backquoted Go string literal, returning the string value
// that s quotes. Unlike the version in strconv, this version
// doesn't mind single quoted strings
func strunquote(s string) (string, error) {
	n := len(s)
	if n < 2 {
		return "", strconv.ErrSyntax
	}
	quote := s[0]
	if quote != s[n-1] {
		return "", strconv.ErrSyntax
	}
	s = s[1 : n-1]

	if quote != '"' && quote != '\'' {
		return "", strconv.ErrSyntax
	}
	if contains(s, '\n') {
		return "", strconv.ErrSyntax
	}

	// Is it trivial? Avoid allocation.
	if !contains(s, '\\') && !contains(s, quote) {
		switch quote {
		case '"':
			return s, nil
		case '\'':
			r, size := utf8.DecodeRuneInString(s)
			if size == len(s) && (r != utf8.RuneError || size != 1) {
				return s, nil
			}
		}
	}

	var runeTmp [utf8.UTFMax]byte
	buf := make([]byte, 0, 3*len(s)/2) // Try to avoid more allocations.
	for len(s) > 0 {
		c, multibyte, ss, err := strconv.UnquoteChar(s, quote)
		if err != nil {
			return "", err
		}
		s = ss
		if c < utf8.RuneSelf || !multibyte {
			buf = append(buf, byte(c))
		} else {
			n := utf8.EncodeRune(runeTmp[:], c)
			buf = append(buf, runeTmp[:n]...)
		}
	}
	return string(buf), nil
}

// contains reports whether the string contains the byte c.
func contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}


