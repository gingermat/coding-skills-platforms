package sort

import "math"

func SortMerge(l []int) {
	length := len(l)

	aux := make([]int, length)
	sortMS(l, aux, 0, length-1)
}

func SortMergeBU(l []int) {
	length := len(l)

	aux := make([]int, length)

	for i := 1; i < length; i *= 2 {
		for lo := 0; lo < length-i; lo += i * 2 {
			mergeMS(l, aux, lo, lo+i-1, int(math.Min(float64(lo+2*i-1), float64(length))))
		}
	}
}

func sortMS(l []int, aux []int, lo int, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2

	sortMS(l, aux, lo, mid)
	sortMS(l, aux, mid+1, hi)

	if l[mid+1] >= l[mid] {
		return
	}

	mergeMS(l, aux, lo, mid, hi)
}

func mergeMS(l []int, aux []int, lo int, mid int, hi int) {
	if !(isSorted(l, lo, mid) && isSorted(l, mid+1, hi)) {
		panic("slice not sorted")
	}

	for k := lo; k <= hi; k++ {
		aux[k] = l[k]
	}

	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			l[k] = aux[j]
			j++
		} else if j > hi {
			l[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			l[k] = aux[j]
			j++
		} else {
			l[k] = aux[i]
			i++
		}
	}

	if !isSorted(l, lo, hi) {
		panic("slice after merge not sorted")
	}
}
