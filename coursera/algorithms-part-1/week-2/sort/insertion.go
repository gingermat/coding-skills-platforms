package sort

func SortInsertion(l []int) {
	var length = len(l)

	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if l[j] < l[j-1] {
				swap(l, j, j-1)
			} else {
				break
			}
		}
	}
}
