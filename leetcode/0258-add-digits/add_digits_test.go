package add_digits

import (
	"testing"
)

var tests = []struct {
	num  int
	want int
}{
	{
		38, 2,
	},
	{
		0, 0,
	},
}

func TestAddDigits(t *testing.T) {
	for _, tt := range tests {
		got := addDigits(tt.num)
		if got != tt.want {
			t.Fatalf("addDigits(%v) = %v, want %v", tt.num, got, tt.want)
		}
	}
}
