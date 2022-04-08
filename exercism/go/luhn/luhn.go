package luhn

import "unicode"

// Valid is given a number determine whether or
//  not it is valid per the Luhn formula
func Valid(input string) bool {
	var (
		isEven   bool
		numCount int
		sum      int
	)

	for i := len(input) - 1; i >= 0; i-- {
		r := rune(input[i])

		if r == ' ' {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}

		i := int(r - '0')

		if isEven {
			i *= 2
			if i > 9 {
				i -= 9
			}
		}
		isEven = !isEven

		sum += i
		numCount++
	}

	if numCount < 2 {
		return false
	}

	return sum%10 == 0
}
