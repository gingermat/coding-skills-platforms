package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	var count []int
	var chars, result []rune

	var curr, prev rune
	for _, chr := range input {
		if prev, curr = curr, chr; curr != prev {
			chars = append(chars, chr)
			count = append(count, 1)
		} else {
			count[len(count)-1]++
		}
	}

	for i, chr := range chars {
		if q := count[i]; q > 1 {
			result = append(result, []rune(strconv.Itoa(q))...)
		}
		result = append(result, chr)
	}

	return string(result)
}

func RunLengthDecode(input string) string {
	var count int
	var result []string

	for _, chr := range input {
		if num, err := strconv.Atoi(string(chr)); err == nil {
			count = count*10 + num
			continue
		}

		if count == 0 {
			count = 1
		}

		result = append(result, strings.Repeat(string(chr), count))
		count = 0
	}

	return strings.Join(result, "")
}
