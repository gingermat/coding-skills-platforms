package darts

import (
	"math"
)

func Score(x, y float64) int {
	radiusLength := math.Hypot(x, y)

	switch {
	case radiusLength <= 1:
		return 10
	case radiusLength <= 5:
		return 5
	case radiusLength <= 10:
		return 1
	}

	return 0
}
