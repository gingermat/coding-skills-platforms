package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var result []string

	if len(rhyme) > 1 {
		for i := 1; i < len(rhyme); i++ {
			chunk := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i])
			result = append(result, chunk)
		}
	}

	if len(rhyme) > 0 {
		chunk := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		result = append(result, chunk)
	}

	return result
}
