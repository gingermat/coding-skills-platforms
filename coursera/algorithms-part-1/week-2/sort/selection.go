package sort

func SortSelection(l []int) {
	length := len(l)

	for i := 0; i < length; i++ {
		min := i

		for j := i + 1; j < length; j++ {
			if l[j] < l[min] {
				min = j
			}
			swap(l, i, min)
		}
	}
}
