package search_insert_position

import (
	"testing"
)

var tests = []struct {
	nums   []int
	target int
	want   int
}{
	{
		[]int{1, 3, 5, 6}, 5, 2,
	},
	{
		[]int{1, 3, 5, 6}, 2, 1,
	},
	{
		[]int{1, 3, 5, 6}, 7, 4,
	},
}

func TestSearchInsert(t *testing.T) {
	for _, tt := range tests {
		got := searchInsert(tt.nums, tt.target)
		if got != tt.want {
			t.Fatalf("searchInsert(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
