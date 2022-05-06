package binary_search

import (
	"testing"
)

var tests = []struct {
	input  []int
	target int
	want   int
}{
	{
		[]int{-1, 0, 3, 5, 9, 12}, 9, 4,
	},
	{
		[]int{-1, 0, 3, 5, 9, 12}, 2, -1,
	},
	{
		[]int{-1, 0, 3, 5, 9, 12}, 13, -1,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, tt := range tests {
		got := search(tt.input, tt.target)
		if got != tt.want {
			t.Fatalf("search(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
