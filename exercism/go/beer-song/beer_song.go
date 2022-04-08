package beer

import "fmt"

func Verses(upper, lower int) (string, error) {
	if upper > 100 || lower < 0 || upper <= lower {
		return "", fmt.Errorf("invalid upper %d or lower %d values", upper, lower)
	}

	var fullText string

	for upper >= lower {
		if verse, err := Verse(upper); err == nil {
			fullText += verse + "\n"
		} else {
			return "", err
		}

		upper--
	}

	return fullText, nil
}

func Verse(bottles int) (string, error) {
	if bottles < 0 || bottles > 100 {
		return "", fmt.Errorf("invalid bottle count %d", bottles)
	}

	if bottles == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\n" +
			"Go to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
	if bottles == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\n" +
			"Take it down and pass it around, no more bottles of beer on the wall.\n", nil
	}
	if bottles == 2 {
		return "2 bottles of beer on the wall, 2 bottles of beer.\n" +
			"Take one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}

	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\n"+
		"Take one down and pass it around, %d bottles of beer on the wall.\n",
		bottles, bottles, bottles-1), nil
}

func Song() string {
	song, _ := Verses(99, 0)
	return song
}
