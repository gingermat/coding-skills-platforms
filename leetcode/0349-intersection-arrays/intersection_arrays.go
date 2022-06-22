package intersection_arrays

func intersection(nums1 []int, nums2 []int) []int {
	var res []int

	nums := make(map[int]struct{})

	for _, num := range nums1 {
		nums[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, ok := nums[num]; ok {
			res = append(res, num)
			delete(nums, num)
		}
	}

	return res
}
