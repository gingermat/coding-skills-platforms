package majority_element

func majorityElement(nums []int) int {
	counts := make(map[int]int)

	var max int

	for _, i := range nums {
		counts[i]++

		if counts[i] > max {
			max = counts[i]
		}
	}

	for _, i := range nums {
		if counts[i] == max {
			return i
		}
	}

	return 0
}
