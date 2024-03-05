package structsmethodsandinterfaces

import (
	"fmt"
	"testing"
)

func TestTriangleArea(t *testing.T) {
	for index, testCase := range []AreaTestCase{
		{Shape: Triangle{Base: 12.0, Height: 6.0}, Expect: 36.0},
		{Shape: Triangle{Base: 1.0, Height: 2.0}, Expect: 2.0},
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

func TestTrianglePerimeter(t *testing.T) {
	for index, testCase := range []PerimeterTestCase{
		{Shape: Triangle{Base: 10.0, Height: 10.0}, Expect: 0.0},
		{Shape: Triangle{Base: 20.0, Height: 15.0}, Expect: 0.0},
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
