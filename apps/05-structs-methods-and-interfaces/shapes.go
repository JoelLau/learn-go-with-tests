package shapes

import "math"

type Shape interface {
	getArea() float64
	getPerimeter() float64
}

type Rectangle struct {
	width  float64
	length float64
}

func (rectangle Rectangle) getPerimeter() float64 {
	return 2.0 * (rectangle.length + rectangle.width)
}

func (rectangle Rectangle) getArea() float64 {
	return rectangle.length * rectangle.width
}

type Circle struct {
	radius float64
}

func (circle Circle) getPerimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func (circle Circle) getArea() float64 {
	return math.Pi * circle.radius * circle.radius
}
