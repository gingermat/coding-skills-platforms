package square_of_sorted_array

func sortedSquares(nums []int) []int {
	length := len(nums)

	l, r := 0, length-1
	answer := make([]int, length)

	for i := length - 1; i >= 0; i-- {
		lv := nums[l] * nums[l]
		rv := nums[r] * nums[r]

		if lv > rv {
			answer[i] = lv
			l++
		} else {
			answer[i] = rv
			r--
		}
	}

	return answer
}
