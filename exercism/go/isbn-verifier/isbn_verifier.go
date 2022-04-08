package isbn

var charToNum = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'X': 10,
}

var maxNumCount int = 10

func IsValidISBN(isbn string) bool {
	var acc, pos int

	for _, chr := range isbn {
		if num, ok := charToNum[chr]; ok {
			if pos >= maxNumCount {
				return false
			}
			if num == 10 && pos != maxNumCount-1 {
				return false
			}

			acc += num * (10 - pos)
			pos++
		}
	}

	if pos < maxNumCount {
		return false
	}

	return acc%11 == 0
}
