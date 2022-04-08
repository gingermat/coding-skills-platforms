// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "unicode"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var (
		skipLetter bool
		result     []rune
	)

	for _, r := range s {
		if !skipLetter && unicode.IsLetter(r) {
			result = append(result, unicode.ToUpper(r))
			skipLetter = true
		}

		if r == ' ' || r == '-' {
			skipLetter = false
		}
	}

	return string(result)
}
