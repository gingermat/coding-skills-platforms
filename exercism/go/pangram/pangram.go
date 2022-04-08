package pangram

func IsPangram(input string) bool {
	var lettersMap int32

	for _, r := range input {
		if 'A' <= r && r <= 'Z' {
			r += 'a' - 'A'
		}

		if 'a' <= r && r <= 'z' {
			lettersMap = lettersMap | 1<<(r-'a')
		}
	}

	return 1<<26-lettersMap == 1
}
