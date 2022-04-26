package sort

func swap(l []int, i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func isSorted(l []int) bool {
	for i := 1; i < len(l); i++ {
		p, q := l[i-1], l[i]

		if p < q {
			return false
		}
	}

	return true
}
