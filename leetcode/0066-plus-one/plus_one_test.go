package plus_one

import (
	"reflect"
	"testing"
)

var tests = []struct {
	digits []int
	want   []int
}{
	{
		[]int{1, 2, 3}, []int{1, 2, 4},
	},
	{
		[]int{4, 3, 2, 1}, []int{4, 3, 2, 2},
	},
	{
		[]int{9}, []int{1, 0},
	},
}

func TestPlusOne(t *testing.T) {
	for _, tt := range tests {
		got := plusOne(tt.digits)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
		}
	}
}
