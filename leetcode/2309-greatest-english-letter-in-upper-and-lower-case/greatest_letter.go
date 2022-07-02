package greatest_letter

func greatestLetter(s string) string {
	letters := make([]rune, 52)

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			letters[r-'a'+26]++
		} else {
			letters[r-'A']++
		}
	}

	for i := len(letters) - 27; i >= 0; i-- {
		if letters[i] > 0 && letters[i+26] > 0 {
			return string(rune(i + 'A'))
		}
	}

	return ""
}
