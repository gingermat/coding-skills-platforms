package cryptosquare

import (
	"math"
	"unicode"
)

func filter(input string) []rune {
	var filtered []rune

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, unicode.ToLower(r))
		}
	}

	return filtered
}

func Encode(input string) string {
	filtered := filter(input)
	filteredLen := len(filtered)

	if filteredLen == 0 {
		return ""
	}

	squareLen := int(math.Ceil(math.Sqrt(float64(filteredLen))))

	var parts [][]rune

	for r := 0; r < squareLen; r++ {
		part := make([]rune, squareLen)

		for c := 0; c < squareLen; c++ {
			pos := r*squareLen + c

			chr := ' '
			if pos < filteredLen {
				chr = filtered[pos]
			} else if c == 0 {
				continue
			}
			part[c] = chr
		}

		if part[0] != '\x00' {
			parts = append(parts, part)
		}
	}

	var result []rune

	for r := 0; r < len(parts[0]); r++ {
		for c := 0; c < len(parts); c++ {
			result = append(result, parts[c][r])
		}

		if r < squareLen-1 {
			result = append(result, ' ')
		}
	}

	return string(result)
}
