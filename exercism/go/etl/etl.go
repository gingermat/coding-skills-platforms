package etl

import (
	"strings"
)

func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)

	for point, m := range input {
		for _, el := range m {
			lowercaseLetter := strings.ToLower(el)
			result[lowercaseLetter] = point
		}
	}
	return result
}
