package twosum

func twoSum(nums []int, target int) []int {
	positionMap := make(map[int]int, len(nums))

	for pos, num := range nums {
		secondTerm := target - num
		if v, ok := positionMap[secondTerm]; ok {
			return []int{v, pos}
		}

		positionMap[num] = pos
	}

	return []int{}
}
