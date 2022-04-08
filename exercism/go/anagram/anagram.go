package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var valid []string

	subject = strings.ToLower(subject)
	subjectMap := make(map[rune]int)

	for _, r := range subject {
		subjectMap[r]++
	}

	for _, c := range candidates {
		if len(c) != len(subject) {
			continue
		}

		candidate := strings.ToLower(c)
		candidateMap := make(map[rune]int)

		for _, r := range candidate {
			candidateMap[r]++
		}

		if reflect.DeepEqual(candidateMap, subjectMap) && candidate != subject {
			valid = append(valid, c)
		}
	}

	return valid
}
