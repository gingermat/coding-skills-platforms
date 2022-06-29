package majority_element

import (
	"testing"
)

var tests = []struct {
	nums []int
	want int
}{
	{
		[]int{3, 2, 3}, 3,
	},
	{
		[]int{2, 2, 1, 1, 1, 2, 2}, 2,
	},
	{
		[]int{1}, 1,
	},
	{
		[]int{8, 8, 7, 7, 7}, 7,
	},
	{
		[]int{6, 6, 6, 7, 7}, 6,
	},
}

func TestMajorityElement(t *testing.T) {
	for _, tt := range tests {
		got := majorityElement(tt.nums)
		if got != tt.want {
			t.Fatalf("majorityElement(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
