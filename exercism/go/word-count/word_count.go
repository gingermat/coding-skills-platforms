package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(s string) Frequency {
	fq := make(Frequency)

	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
	})

	for _, w := range words {
		if w[0] == '\'' {
			w = strings.Trim(w, "'")
		}

		fq[strings.ToLower(w)]++
	}

	return fq
}
