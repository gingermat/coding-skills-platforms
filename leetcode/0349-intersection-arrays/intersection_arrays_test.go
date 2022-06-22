package intersection_arrays

import (
	"reflect"
	"sort"
	"testing"
)

var tests = []struct {
	nums1 []int
	nums2 []int
	want  []int
}{
	{
		[]int{1, 2, 2, 1}, []int{2, 2}, []int{2},
	},
	{
		[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9},
	},
}

func TestIntersection(t *testing.T) {
	for _, tt := range tests {
		got := intersection(tt.nums1, tt.nums2)

		sort.Ints(got)
		sort.Ints(tt.want)

		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("intersection(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
		}
	}
}
