package scrabble

import (
	"strings"
	"unicode"
)

var scoresMap = map[rune]int{
	rune('A'): 1,
	rune('E'): 1,
	rune('I'): 1,
	rune('O'): 1,
	rune('U'): 1,
	rune('L'): 1,
	rune('N'): 1,
	rune('R'): 1,
	rune('S'): 1,
	rune('T'): 1,
	rune('D'): 2,
	rune('G'): 2,
	rune('B'): 3,
	rune('C'): 3,
	rune('M'): 3,
	rune('P'): 3,
	rune('F'): 4,
	rune('H'): 4,
	rune('V'): 4,
	rune('W'): 4,
	rune('Y'): 4,
	rune('K'): 5,
	rune('J'): 8,
	rune('X'): 8,
	rune('Q'): 10,
	rune('Z'): 10,
}

// Score compute the scrabble score for input word
func Score(input string) int {
	var result int

	for _, ch := range strings.Split(input, "") {
		r := rune(ch[0])
		r = unicode.ToUpper(r)
		if score, found := scoresMap[r]; found {
			result += score
		}
	}

	return result
}
