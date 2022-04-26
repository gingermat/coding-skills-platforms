package sort

import "fmt"

type Rotation int

const (
	RotationClockwise Rotation = iota - 1
	RotationCounterClockwise
	RotationsCollinear
)

type Point2D struct {
	X int
	Y int
}

func (p Point2D) String() string {
	return fmt.Sprintf("Point2D(x=%d, y=%d)", p.X, p.Y)
}

func NewPoint2D(x, y int) *Point2D {
	return &Point2D{
		X: x,
		Y: y,
	}
}

func CWW(a, b, c *Point2D) Rotation {
	area := (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)

	switch {
	case area < 0:
		return RotationClockwise
	case area > 0:
		return RotationCounterClockwise
	default:
		return RotationsCollinear
	}
}
