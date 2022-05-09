package square_of_sorted_array

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input []int
	want  []int
}{
	{
		[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100},
	},
	{
		[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121},
	},
	{
		[]int{0}, []int{0},
	},
	{
		[]int{-3, 3}, []int{9, 9},
	},
	{
		[]int{-3, 3, 3}, []int{9, 9, 9},
	},
	{
		[]int{-5, -3, -2, -1}, []int{1, 4, 9, 25},
	},
}

func TestSortedSquares(t *testing.T) {
	for _, tt := range tests {
		got := sortedSquares(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("sortedSquares(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
