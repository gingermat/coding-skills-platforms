package rotate_array

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums []int
	k    int
	want []int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4},
	},
	{
		[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100},
	},
}

func TestRotateArray(t *testing.T) {
	for _, tt := range tests {
		got := tt.nums[:]
		rotate(got, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("rotate(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
