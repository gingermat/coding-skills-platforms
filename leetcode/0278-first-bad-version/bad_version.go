package first_bad_version

func firstBadVersion(n int) int {
	min, max := 1, n

	for min < max {
		mid := min + (max-min)/2

		if isBadVersion(mid) {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return min
}
