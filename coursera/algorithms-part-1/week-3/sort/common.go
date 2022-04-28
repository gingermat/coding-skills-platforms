package sort

import "math/rand"

func swap(l []int, i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func shuffle(l []int) {
	rand.Shuffle(len(l), func(i, j int) {
		l[i], l[j] = l[j], l[i]
	})
}

func isSorted(l []int, j int, k int) bool {
	for i := k; i > j; i-- {
		if l[i] < l[i-1] {
			return false
		}
	}
	return true
}
