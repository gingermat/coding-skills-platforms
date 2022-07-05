package license_key

func licenseKeyFormatting(s string, k int) string {
	var res []rune

	for i, c := len(s)-1, 0; i >= 0; i-- {
		char := s[i]

		if char == '-' {
			continue
		}

		if char >= 'a' && char <= 'z' {
			char -= 'a' - 'A'
		}

		if c >= k {
			res = append(res, '-')
			c = 0
		}

		res = append(res, rune(char))
		c++
	}

	for p, q := 0, len(res)-1; p < q; p, q = p+1, q-1 {
		res[p], res[q] = res[q], res[p]
	}

	return string(res)
}
