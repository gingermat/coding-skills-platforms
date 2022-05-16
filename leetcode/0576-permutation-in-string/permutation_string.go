package permutation_string

const letterCount = 26

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// make map alternative with letter counts
	s1Map := [letterCount]int{}
	for _, r := range s1 {
		s1Map[r-'a']++
	}

	for i, r := range s2 {
		// skip if symbol not exists in s1
		if idx := r - 'a'; s1Map[idx] == 0 {
			continue
		}

		// check left substring
		if lPos := i - len(s1) + 1; lPos >= 0 {
			if compare(s2[lPos:i+1], s1Map) {
				return true
			}
		}

		// check right substring
		if rPos := i + len(s1); rPos <= len(s2) {
			if compare(s2[i:rPos], s1Map) {
				return true
			}
		}
	}

	return false
}

func compare(s string, m [letterCount]int) bool {
	s2MapCopy := m[:]

	for _, r := range s {
		idx := r - 'a'
		s2MapCopy[idx]--

		if s2MapCopy[idx] < 0 {
			return false
		}
	}

	return true
}
