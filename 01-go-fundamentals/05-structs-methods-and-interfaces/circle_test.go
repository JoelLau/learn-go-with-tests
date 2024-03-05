package structsmethodsandinterfaces

import (
	"fmt"
	"testing"
)

func TestCircleArea(t *testing.T) {
	for index, testCase := range []AreaTestCase{
		{Shape: Circle{Radius: 10}, Expect: 314.159263589793},
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

func TestCirclePerimeter(t *testing.T) {
	for index, testCase := range []PerimeterTestCase{
		{Shape: Circle{Radius: 10}, Expect: 314.16},
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
