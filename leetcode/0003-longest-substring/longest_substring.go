package longest_substring

func lengthOfLongestSubstring(s string) int {
	alreadySeen := make(map[rune]int)

	var maxLength, j int

	for i, r := range s {
		if lastSeen, ok := alreadySeen[r]; ok && lastSeen > j {
			j = lastSeen
		}

		if partLength := i - j + 1; partLength > maxLength {
			maxLength = partLength
		}

		alreadySeen[r] = i + 1
	}

	return maxLength
}
