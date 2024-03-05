package structsmethodsandinterfaces

import "math"

type Circle struct {
	Radius float64
}

var _ Shape = Circle{}

const pi = math.Pi

func (c Circle) Area() float64 {
	return pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * pi * c.Radius
}
