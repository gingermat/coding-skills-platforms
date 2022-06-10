package is_subsequence

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	var si int

	for ti := 0; si < len(s) && ti < len(t); ti++ {
		if s[si] == t[ti] {
			si++
		}
	}

	return si == len(s)
}
