package two_sum_array

import (
	"reflect"
	"testing"
)

var tests = []struct {
	numbers []int
	target  int
	want    []int
}{
	{
		[]int{2, 7, 11, 15}, 9, []int{1, 2},
	},
	{
		[]int{2, 3, 4}, 6, []int{1, 3},
	},
	{
		[]int{-1, 0}, -1, []int{1, 2},
	},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range tests {
		got := twoSum(tt.numbers, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("twoSum(%v, %v) = %v, want %v", tt.numbers, tt.target, got, tt.want)
		}
	}
}
