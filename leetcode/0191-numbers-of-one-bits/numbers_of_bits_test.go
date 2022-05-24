package number_of_bits

import (
	"strconv"
	"testing"
)

var tests = []struct {
	n    string
	want int
}{
	{
		"00000000000000000000000000001011", 3,
	},
	{
		"0000000000000000000000010000000", 1,
	},
}

func TestHammingWeight(t *testing.T) {
	for _, tt := range tests {
		n, _ := strconv.ParseUint(tt.n, 2, 32)

		got := hammingWeight(uint32(n))
		if got != tt.want {
			t.Fatalf("isPowerOfTwo(%v) = %v, want %v", n, got, tt.want)
		}
	}
}
