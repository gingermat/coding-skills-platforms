package rotationalcipher

func RotationalCipher(text string, shift int) string {
	var result []rune

	for _, r := range text {
		if 'a' <= r && r <= 'z' {
			r = (r-'a'+int32(shift))%26 + 'a'
		} else if 'A' <= r && r <= 'Z' {
			r = (r-'A'+int32(shift))%26 + 'A'
		}
		result = append(result, r)
	}

	return string(result)
}
