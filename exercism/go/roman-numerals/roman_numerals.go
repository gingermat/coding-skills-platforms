package romannumerals

import (
	"fmt"
	"strings"
)

var (
	dividers      = [...]int{1000, 500, 100, 50, 10, 5, 1}
	dividersChars = [...]string{"M", "D", "C", "L", "X", "V", "I"}
)

func ToRomanNumeral(num int) (string, error) {
	if num <= 0 {
		return "", fmt.Errorf("number `%v` must be > 0", num)
	}
	if num > 3000 {
		return "", fmt.Errorf("number `%v` larger than about 3000", num)
	}

	var s strings.Builder

	for i, divider := range dividers {
		if num <= 0 {
			break
		}

		if dr := num / divider; dr > 0 {
			chunks := strings.Repeat(dividersChars[i], dr)
			s.WriteString(chunks)

			num -= divider * dr
		}
	}

	return s.String(), nil
}
