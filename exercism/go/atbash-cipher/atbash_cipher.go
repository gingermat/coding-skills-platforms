package atbash

func Atbash(input string) string {
	var (
		counter int
		output  []rune
	)

	for _, chr := range input {
		switch {
		case 'A' <= chr && chr <= 'Z':
			chr += 'a' - 'A'
			fallthrough
		case 'a' <= chr && chr <= 'z':
			chr = 'a' + 'z' - chr
		case '0' <= chr && chr <= '9':
			break
		default:
			continue
		}

		if counter%5 == 0 && counter > 0 {
			output = append(output, ' ')
		}

		output = append(output, chr)
		counter++
	}

	return string(output)
}
