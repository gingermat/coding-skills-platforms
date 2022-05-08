package search_insert_position

func searchInsert(nums []int, target int) int {
	min, max := 0, len(nums)-1

	for min <= max {
		mid := min + (max-min)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return min
}
