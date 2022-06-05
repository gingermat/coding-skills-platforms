package valid_parentheses

func isValid(s string) bool {
	OpenToClose := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	var stack []rune

	for _, r := range s {
		length := len(stack)
		c, ok := OpenToClose[r]

		if length == 0 && !ok {
			return false
		}
		if ok {
			stack = append(stack, c)
			continue
		}
		if stack[length-1] != r {
			return false
		}

		stack = stack[:length-1]
	}

	return len(stack) == 0
}
