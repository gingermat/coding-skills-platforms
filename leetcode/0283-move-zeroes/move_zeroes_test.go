package move_zeroes

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums []int
	want []int
}{
	{
		[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0},
	},
	{
		[]int{0}, []int{0},
	},
}

func TestMoveZeroes(t *testing.T) {
	for _, tt := range tests {
		got := tt.nums[:]
		moveZeroes(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("moveZeroes(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
