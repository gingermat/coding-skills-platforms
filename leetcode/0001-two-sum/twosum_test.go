package twosum

import (
	"reflect"
	"testing"
)

var tests = []struct {
	num    []int
	target int
	want   []int
}{
	{
		[]int{2, 7, 11, 15}, 9, []int{0, 1},
	},
	{
		[]int{3, 2, 4}, 6, []int{1, 2},
	},
	{
		[]int{3, 3}, 6, []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range tests {
		got := twoSum(tt.num, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("twoSum(%v, %d) = %v, want %v", tt.num, tt.target, got, tt.want)
		}
	}
}
