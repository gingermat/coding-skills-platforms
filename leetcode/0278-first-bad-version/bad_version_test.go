package first_bad_version

import (
	"testing"
)

var isBadVersion func(int) bool

var tests = []struct {
	input    int
	mockFunc func(int) bool
	want     int
}{
	{
		5, func(version int) bool {
			return version > 3
		},
		4,
	},
	{
		1, func(version int) bool {
			return version > 1
		},
		1,
	},
	{
		2,
		func(version int) bool {
			return version > 1
		},
		2,
	},
	{
		1000000, func(version int) bool {
			return version > 72
		},
		73,
	},
}

func TestFirstBadVersion(t *testing.T) {
	for _, tt := range tests {
		isBadVersion = tt.mockFunc

		got := firstBadVersion(tt.input)
		if got != tt.want {
			t.Fatalf("firstBadVersion(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
