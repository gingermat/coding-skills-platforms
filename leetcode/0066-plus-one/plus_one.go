package plus_one

func plusOne(digits []int) []int {
	term := 1

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += term

		if term == 0 {
			return digits
		}

		if digits[i] > 9 {
			digits[i] -= 10
		} else {
			term = 0
		}

		if term == 1 && i == 0 {
			digits = append([]int{1}, digits...)
		}
	}

	return digits
}
