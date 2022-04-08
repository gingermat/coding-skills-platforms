package armstrong

import (
	"fmt"
	"math"
)

func IsNumber(num int) bool {
	var acc int

	numAsString := fmt.Sprintf("%v", num)
	numLength := len(numAsString)

	for _, r := range numAsString {
		acc += int(math.Pow(float64(r-'0'), float64(numLength)))
	}

	return acc == num
}
