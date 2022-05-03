package backspace_compare

func backspaceCompare(s string, t string) bool {
	clean := func(s string) string {
		var items []rune

		for _, r := range s {
			if r != '#' {
				items = append(items, r)
			} else if len(items) > 0 {
				items = items[:len(items)-1]
			}
		}

		return string(items)
	}

	return clean(s) == clean(t)
}
