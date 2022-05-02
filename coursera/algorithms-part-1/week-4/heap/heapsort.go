package heap

func SortHeap(l []int) {
	n := len(l)

	for k := n / 2; k >= 1; k-- {
		sink(l, k, n)
	}

	for n > 1 {
		exch(l, 1, n)
		n--

		sink(l, 1, n)
	}
}

func sink(l []int, k int, n int) {
	for 2*k <= n {
		j := 2 * k

		if j < n && less(l, j, j+1) {
			j++
		}

		if !less(l, k, j) {
			break
		}

		exch(l, k, j)
		k = j
	}
}

func exch(l []int, i int, j int) {
	l[i-1], l[j-1] = l[j-1], l[i-1]
}

func less(l []int, i int, j int) bool {
	return l[i-1] < l[j-1]
}
