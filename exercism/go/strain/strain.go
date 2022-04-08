package strain

type Ints []int
type Lists []Ints
type Strings []string

func (i Ints) Keep(f func(int) bool) Ints {
	var r Ints

	for _, el := range i {
		if f(el) {
			r = append(r, el)
		}
	}

	return r
}

func (i Ints) Discard(f func(int) bool) Ints {
	return i.Keep(func(i int) bool {
		return !f(i)
	})
}

func (l Lists) Keep(f func([]int) bool) Lists {
	var r Lists

	for _, el := range l {
		if f(el) {
			r = append(r, el)
		}
	}

	return r
}

func (s Strings) Keep(f func(string) bool) Strings {
	var r Strings

	for _, el := range s {
		if f(el) {
			r = append(r, el)
		}
	}

	return r
}
