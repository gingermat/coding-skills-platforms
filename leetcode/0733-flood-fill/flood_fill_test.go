package flood_fill

import (
	"reflect"
	"testing"
)

var tests = []struct {
	image    [][]int
	sr       int
	sc       int
	newColor int
	want     [][]int
}{
	{
		[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		1,
		1,
		2,
		[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
	},
	{
		[][]int{{0, 0, 0}, {0, 0, 0}},
		0,
		0,
		2,
		[][]int{{2, 2, 2}, {2, 2, 2}},
	},
}

func TestFloodFill(t *testing.T) {
	for _, tt := range tests {
		got := floodFill(tt.image, tt.sr, tt.sc, tt.newColor)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("floodFill(%v, %v, %v, %v) = %v, want %v",
				tt.image, tt.sr, tt.sc, tt.newColor, got, tt.want)
		}
	}
}
