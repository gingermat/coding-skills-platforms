package sort

func SortQuick(l []int) {
	shuffle(l)

	sortQ(l, 0, len(l)-1)
}

func sortQ(l []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := partition(l, lo, hi)

	sortQ(l, lo, j-1)
	sortQ(l, j+1, hi)
}

func partition(l []int, lo int, hi int) int {
	i, j := lo, hi+1

	for {
		for i += 1; i != hi && l[i] < l[lo]; i++ {
		}
		for j -= 1; j != lo && l[lo] < l[j]; j-- {
		}

		if i >= j {
			break
		}
		swap(l, i, j)
	}

	swap(l, lo, j)
	return j
}
