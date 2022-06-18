package ugly_number

func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}

	res := []int{1}

	var a, b, c int

	for len(res) < n {
		m2, m3, m5 := res[a]*2, res[b]*3, res[c]*5

		min := minInt(m2, minInt(m3, m5))

		res = append(res, min)

		if m2 == min {
			a++
		}

		if m3 == min {
			b++
		}

		if m5 == min {
			c++
		}
	}

	return res[len(res)-1]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}
