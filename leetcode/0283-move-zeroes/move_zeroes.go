package move_zeroes

func moveZeroes(nums []int) {
	var pos int

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos++
		}
	}
}
