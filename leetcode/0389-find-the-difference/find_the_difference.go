package find_the_difference

const LettersCounts = 26

func findTheDifference(s string, t string) byte {
	counts := make([]int, LettersCounts)

	for _, r := range s {
		counts[r-'a']++
	}

	for _, r := range t {
		counts[r-'a']--

		if counts[r-'a'] < 0 {
			return byte(r)
		}
	}

	return 0
}
