package sort

func SortShell(l []int) {
	length := len(l)
	h := 1

	for ; h < length/3; h = 3*h + 1 {
	}

	for ; h > 1; h /= 3 {
		for h >= 1 {
			SortInsertion(l[h-1 : length])
			h /= 3
		}
	}
}
