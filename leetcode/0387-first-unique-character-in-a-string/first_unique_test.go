package first_unique

import (
	"testing"
)

var tests = []struct {
	s    string
	want int
}{
	{
		"leetcode", 0,
	},
	{
		"loveleetcode", 2,
	},
	{
		"aabb", -1,
	},
}

func TestFirstUniqChar(t *testing.T) {
	for _, tt := range tests {
		got := firstUniqChar(tt.s)
		if got != tt.want {
			t.Fatalf("firstUniqChar(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
