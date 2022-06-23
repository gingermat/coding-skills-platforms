package nim_game

import "testing"

var tests = []struct {
	n    int
	want bool
}{
	{
		4, false,
	},
	{
		2, true,
	},
	{
		1, true,
	},
	{
		8, false,
	},
	{
		7, true,
	},
}

func TestCanWinNim(t *testing.T) {
	for _, tt := range tests {
		got := canWinNim(tt.n)
		if got != tt.want {
			t.Fatalf("canWinNim(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
