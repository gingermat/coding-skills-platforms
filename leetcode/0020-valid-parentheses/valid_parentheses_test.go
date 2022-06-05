package valid_parentheses

import (
	"testing"
)

var tests = []struct {
	s    string
	want bool
}{
	{
		"()", true,
	},
	{
		"()[]{}", true,
	},
	{
		"(]", false,
	},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := isValid(tt.s)
		if got != tt.want {
			t.Fatalf("isValid(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
