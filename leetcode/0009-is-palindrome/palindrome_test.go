package palindrome

import (
	"testing"
)

var tests = []struct {
	input int
	want  bool
}{
	{
		121, true,
	},
	{
		-121, false,
	},
	{
		10, false,
	},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range tests {
		got := isPalindrome(tt.input)
		if got != tt.want {
			t.Fatalf("isPalindrome(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
