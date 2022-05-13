package reverse_string

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input []byte
	want  []byte
}{
	{
		[]byte("hello"), []byte("olleh"),
	},
	{
		[]byte("Hannah"), []byte("hannaH"),
	},
}

func TestReverseString(t *testing.T) {
	for _, tt := range tests {
		input := tt.input[:]

		reverseString(input)
		if !reflect.DeepEqual(input, tt.want) {
			t.Fatalf("reverseString(%s) = %s, want %s", tt.input, input, tt.want)
		}
	}
}
