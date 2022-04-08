// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "invalid"     // not a triangle
	Equ Kind = "equilateral" // equilateral
	Iso Kind = "isosceles"   // isosceles
	Sca Kind = "scalene"     // scalene
)

// KindFromSides determine if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case a <= 0 || b <= 0 || c <= 0:
		return NaT
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return NaT
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return NaT
	case a+b < c || b+c < a || c+a < b:
		return NaT
	case a == b && b == c && c == a:
		return Equ
	case (a == b && b != c) || (b == c && c != a) || (c == a && a != b):
		return Iso
	case a != b && b != c && c != a:
		return Sca
	}

	return NaT
}
