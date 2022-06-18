package ugly_number

import (
	"testing"
)

var tests = []struct {
	n    int
	want int
}{
	{
		10, 12,
	},
	{
		1, 1,
	},
	{
		11, 15,
	},
}

func TestNthUglyNumber(t *testing.T) {
	for _, tt := range tests {
		got := nthUglyNumber(tt.n)
		if got != tt.want {
			t.Fatalf("nthUglyNumber(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
