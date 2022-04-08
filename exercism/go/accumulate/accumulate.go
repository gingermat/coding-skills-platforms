package accumulate

func Accumulate(els []string, op func(string) string) []string {
	result := make([]string, len(els))

	for i, el := range els {
		result[i] = op(el)
	}

	return result
}
