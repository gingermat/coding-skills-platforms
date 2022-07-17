package maximum_product

func maxProduct(nums []int) int {
	var a, b int

	for _, v := range nums {
		if v > b {
			b = v
		}

		if b > a {
			a, b = b, a
		}
	}

	return (a - 1) * (b - 1)
}
