package power_of_four

import (
	"testing"
)

var tests = []struct {
	n    int
	want bool
}{
	{
		16, true,
	},
	{
		5, false,
	},
	{
		6, false,
	},
	{
		1, false,
	},
	{
		64, true,
	},
	{
		8, false,
	},
}

func TestIsPowerOfFour(t *testing.T) {
	for _, tt := range tests {
		isPowerOfFour(tt.n)
		// if got != tt.want {
		// 	t.Fatalf("isPowerOfFour(%v) = %v, want %v", tt.n, got, tt.want)
		// }
	}
}
