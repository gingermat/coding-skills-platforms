package power_of_two

import (
	"testing"
)

var tests = []struct {
	n    int
	want bool
}{
	{
		1, true,
	},
	{
		16, true,
	},
	{
		3, false,
	},
}

func TestIsPowerOfTwo(t *testing.T) {
	for _, tt := range tests {
		got := isPowerOfTwo(tt.n)
		if got != tt.want {
			t.Fatalf("isPowerOfTwo(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
