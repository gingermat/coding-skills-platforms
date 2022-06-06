package power_of_three

import (
	"testing"
)

var tests = []struct {
	n    int
	want bool
}{
	{
		27, true,
	},
	{
		0, false,
	},
	{
		9, true,
	},
}

func TestIsPowerOfThree(t *testing.T) {
	for _, tt := range tests {
		got := isPowerOfThree(tt.n)
		if got != tt.want {
			t.Fatalf("isPowerOfThree(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
