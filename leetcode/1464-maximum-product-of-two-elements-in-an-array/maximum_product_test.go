package maximum_product

import (
	"testing"
)

var tests = []struct {
	input []int
	want  int
}{
	{
		[]int{3, 4, 5, 2}, 12,
	},
	{
		[]int{1, 5, 4, 5}, 16,
	},
	{
		[]int{3, 7}, 12,
	},
}

func TestMaxProduct(t *testing.T) {
	for _, tt := range tests {
		got := maxProduct(tt.input)
		if got != tt.want {
			t.Fatalf("maxProduct(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
