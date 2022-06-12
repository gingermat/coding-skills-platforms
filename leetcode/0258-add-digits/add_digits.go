package add_digits

import "fmt"

func addDigits(num int) int {
	for num >= 10 {
		var sum int

		for num > 0 {
			sum += num % 10
			num = num / 10
		}

		num = sum
	}

	return num
}
