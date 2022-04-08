package stringset

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func (s Set) String() string {
	var keys []string

	for k := range s {
		keys = append(keys, fmt.Sprintf("%q", k))
	}

	return "{" + strings.Join(keys, ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(el string) bool {
	return s[el]
}

func (s Set) Add(el string) {
	s[el] = true
}

func New() Set {
	return NewFromSlice(
		[]string{},
	)
}

func NewFromSlice(slice []string) Set {
	s := make(Set)

	for _, e := range slice {
		s[e] = true
	}

	return s
}

func Subset(s1, s2 Set) bool {
	return Difference(s1, s2).IsEmpty()
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	long, short := s1, s2

	if len(long) < len(short) {
		long, short = short, long
	}

	var slice []string

	for k := range long {
		if short.Has(k) {
			slice = append(slice, k)
		}
	}

	return NewFromSlice(slice)
}

func Difference(s1, s2 Set) Set {
	newSet := New()

	for k := range s1 {
		if !s2.Has(k) {
			newSet.Add(k)
		}
	}

	return newSet
}

func Union(s1, s2 Set) Set {
	newSet := New()

	for _, s := range []Set{s1, s2} {
		for k := range s {
			newSet.Add(k)
		}
	}

	return newSet
}
