package fizz_buzz

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	var res []string

	for i := 1; i <= n; i++ {
		var value string

		switch {
		case i%15 == 0:
			value = "FizzBuzz"
		case i%5 == 0:
			value = "Buzz"
		case i%3 == 0:
			value = "Fizz"
		default:
			value = strconv.Itoa(i)
		}

		res = append(res, value)
	}

	return res
}
