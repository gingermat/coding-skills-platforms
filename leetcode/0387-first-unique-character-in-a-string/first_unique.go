package first_unique

func firstUniqChar(s string) int {
	found := make(map[rune]int)

	for i, r := range s {
		if _, ok := found[r]; !ok {
			found[r] = i
		} else {
			found[r] = -1
		}
	}

	result := -1

	for _, v := range found {
		if v == -1 {
			continue
		}

		if result == -1 || v < result {
			result = v
		}
	}

	return result
}
