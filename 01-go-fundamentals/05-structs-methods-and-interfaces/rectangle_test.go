package structsmethodsandinterfaces

import (
	"fmt"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	for index, testCase := range []AreaTestCase{
		{Shape: Rectangle{Width: 12.0, Height: 6.0}, Expect: 72.0},
		{Shape: Rectangle{Width: 1.0, Height: 2.0}, Expect: 2.0},
	} {
		name := fmt.Sprintf("Test %q", index)
		t.Run(name, func(t *testing.T) {
			have := testCase.Shape.Area()
			want := testCase.Expect

			if have != want {
				t.Errorf("%#v have %g want %g", testCase.Shape, have, want)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	for index, testCase := range []PerimeterTestCase{
		{Shape: Rectangle{Width: 10.0, Height: 10.0}, Expect: 40.0},
		{Shape: Rectangle{Width: 20.0, Height: 15.0}, Expect: 70.0},
	} {
		name := fmt.Sprintf("Test %q", index)
		t.Run(name, func(t *testing.T) {
			have := testCase.Shape.Perimeter()
			want := testCase.Expect

			if have != want {
				t.Errorf("%#v have %g want %g", testCase.Shape, have, want)
			}
		})
	}
}
