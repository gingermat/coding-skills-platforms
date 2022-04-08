package isogram

import "unicode"

// IsIsogram check input string for isogram
func IsIsogram(input string) bool {
	charMap := make(map[rune]bool)

	if input == "" {
		return true
	}

	for _, ch := range input {
		r := rune(ch)

		if !unicode.IsLetter(r) {
			continue
		}

		r = unicode.ToLower(r)

		if _, found := charMap[r]; found {
			return false
		}

		charMap[r] = true
	}
	return true
}
