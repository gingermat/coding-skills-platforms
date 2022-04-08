package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var diffCounter int

	if len(a) != len(b) {
		return diffCounter, errors.New("Uncomparable strings")
	}

	for idx := 0; idx < len(a); idx++ {
		if a[idx] != b[idx] {
			diffCounter++
		}
	}

	return diffCounter, nil

}
