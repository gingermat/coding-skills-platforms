package valid_anagram

import (
	"testing"
)

var tests = []struct {
	s    string
	t    string
	want bool
}{
	{
		"anagram",
		"nagaram",
		true,
	},
	{
		"rat",
		"car",
		false,
	},
}

func TestIsAnagram(t *testing.T) {
	for _, tt := range tests {
		got := isAnagram(tt.s, tt.t)
		if got != tt.want {
			t.Fatalf("isAnagram(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
