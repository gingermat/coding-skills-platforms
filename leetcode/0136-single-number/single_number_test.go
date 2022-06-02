package single_number

import (
	"testing"
)

var tests = []struct {
	nums []int
	want int
}{
	{
		[]int{2, 2, 1}, 1,
	},
	{
		[]int{4, 1, 2, 1, 2}, 4,
	},
	{
		[]int{1}, 1,
	},
}

func TestSingleNumber(t *testing.T) {
	for _, tt := range tests {
		got := singleNumber(tt.nums)
		if got != tt.want {
			t.Fatalf("singleNumber(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
