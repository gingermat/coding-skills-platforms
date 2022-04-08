package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts concurently the frequency of each rune in a given text
// and returns this data as a FreqMap
func ConcurrentFrequency(strings []string) FreqMap {
	var ch = make(chan FreqMap, 10)

	m := FreqMap{}

	for _, s := range strings {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	for range strings {
		for r, c := range <-ch {
			m[r] += c
		}
	}

	return m
}
