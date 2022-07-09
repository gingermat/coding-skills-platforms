package fizz_buzz

import (
	"reflect"
	"testing"
)

var tests = []struct {
	n    int
	want []string
}{
	{
		3,
		[]string{"1", "2", "Fizz"},
	},
	{
		5,
		[]string{"1", "2", "Fizz", "4", "Buzz"},
	},
	{
		15,
		[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
	},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range tests {
		got := fizzBuzz(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("fizzBuzz(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
