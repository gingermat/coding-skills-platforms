package is_subsequence

import (
	"testing"
)

var tests = []struct {
	s    string
	t    string
	want bool
}{
	{
		"abc", "ahbgdc", true,
	},
	{
		"axc", "ahbgdc", false,
	},
	{
		"", "ahbgdc", true,
	},
	{
		"aaaaaa", "bbaaaa", false,
	},
	{
		"bb", "ahbgdc", false,
	},
}

func TestIsSubsequence(t *testing.T) {
	for _, tt := range tests {
		got := isSubsequence(tt.s, tt.t)
		if got != tt.want {
			t.Fatalf("isSubsequence(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
