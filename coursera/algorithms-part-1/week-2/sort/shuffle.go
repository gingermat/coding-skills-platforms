package sort

import (
	"math/rand"
)

func Shuffle(l []int) {
	length := len(l)

	for i := 0; i < length; i++ {
		j := rand.Intn(i + 1)
		swap(l, i, j)
	}
}
