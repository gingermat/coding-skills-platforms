package permutation_string

import (
	"testing"
)

var tests = []struct {
	s1   string
	s2   string
	want bool
}{
	{
		"ab", "eidbaooo", true,
	},
	{
		"ab", "eidboaoo", false,
	},
}

func TestCheckInclusion(t *testing.T) {
	for _, tt := range tests {
		got := checkInclusion(tt.s1, tt.s2)
		if got != tt.want {
			t.Fatalf("checkInclusion(%v, %v) = %v, want %v", tt.s1, tt.s2, got, tt.want)
		}
	}
}
