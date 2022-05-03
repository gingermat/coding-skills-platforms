package backspace_compare

import "testing"

var tests = []struct {
	s    string
	t    string
	want bool
}{
	{
		"ab#c", "ad#c", true,
	},
	{
		"ab##", "c#d#", true,
	},
	{
		"a#c", "b", false,
	},
}

func TestBackspaceCompare(t *testing.T) {
	for _, tt := range tests {
		got := backspaceCompare(tt.s, tt.t)
		if got != tt.want {
			t.Fatalf("backspaceCompare(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
