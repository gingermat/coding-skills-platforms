package two_sum_array

func twoSum(numbers []int, target int) []int {
	for l, r := 1, len(numbers); l < r; r-- {
		for ll := l; ll < r; ll++ {
			summ := numbers[ll-1] + numbers[r-1]

			if summ > target {
				break
			}

			if summ == target {
				return []int{ll, r}
			}
		}
	}
	return []int{}
}
