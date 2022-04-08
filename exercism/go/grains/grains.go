package grains

import "fmt"

// Square is calculate 2**n
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("bad square number %d: must be 1-64", n)
	}

	return 1 << (n - 1), nil
}

// Total is calculate the number of grains of wheat on a chessboard
func Total() uint64 {
	return 1<<64 - 1
}
