package valid_anagram

const LettersCounts = 26

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make([]int, LettersCounts)

	for i := 0; i < len(s); i++ {
		if sv, tv := s[i], t[i]; sv == tv {
			continue
		} else {
			counts[sv-'a']++
			counts[tv-'a']--
		}
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}
