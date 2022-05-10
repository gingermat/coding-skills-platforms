package rotate_array

func rotate(nums []int, k int) {
	k %= len(nums)

	if k == 0 {
		return
	} else if k < 0 {
		k += 7
	}

	lastIdx := len(nums) - 1

	for i := 0; i < k; i++ {
		tmp := nums[lastIdx]

		for l := lastIdx; l > 0; l-- {
			nums[l] = nums[l-1]
		}

		nums[0] = tmp
	}
}
