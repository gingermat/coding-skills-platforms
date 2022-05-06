package binary_search

func search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		p := (l + r) / 2

		if nums[p] == target {
			return p
		}

		if nums[p] > target {
			r = p - 1
		} else {
			l = p + 1
		}

	}
	return -1
}
