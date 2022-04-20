package longest_substring

import (
	"testing"
)

var tests = []struct {
	input    string
	expected int
}{
	{
		"abcabcbb", 3,
	},
	{
		"bbbbb", 1,
	},
	{
		"pwwkew", 3,
	},
	{
		" ", 1,
	},
	{
		"au", 2,
	},
	{
		"aab", 2,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.input)
		if got != tt.expected {
			t.Fatalf("lengthOfLongestSubstring(%v) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}
