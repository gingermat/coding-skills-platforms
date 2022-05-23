package reverse_bits

import (
	"testing"
)

var tests = []struct {
	num  uint32
	want uint32
}{
	{
		43261596, 964176192,
	},
	{
		4294967293, 3221225471,
	},
}

func TestReverseBits(t *testing.T) {
	for _, tt := range tests {
		got := reverseBits(tt.num)
		if got != tt.want {
			t.Fatalf("reverseBits(%v) = %v, want %v", tt.num, got, tt.want)
		}
	}
}
