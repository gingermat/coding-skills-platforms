package prime

import (
	"math"
)

func Nth(value int) (int, bool) {
	if value == 0 {
		return 0, false
	}

	n := 1

	for i := 1; ; i++ {
		if isPrime(i) {
			n++
		}

		if n > value {
			return i, true
		}
	}
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
