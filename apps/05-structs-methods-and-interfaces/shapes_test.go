package shapes

import "testing"
import "fmt"

type PerimeterTestCase struct {
	shape     Shape
	perimeter float64
}

type AreaTestCase struct {
	shape Shape
	area  float64
}

func float64Equals(t *testing.T, a, b float64) {
	t.Helper()

	if a != b {
		t.Errorf("Expecting %v and %v to match", a, b)
	}
}

func TestRectangle(t *testing.T) {
	t.Run("implements interface: Shape", func(t *testing.T) {
		var _ Shape = Rectangle{1.0, 1.0}
	})

	for _, perimeterTestCase := range []PerimeterTestCase{
		{shape: Rectangle{6.0, 2.0}, perimeter: 16.0},
		{shape: Rectangle{1.1, 1.1}, perimeter: 4.4},
	} {
		rectangle := perimeterTestCase.shape.(Rectangle)
		t.Run(
			fmt.Sprintf("should return %v when given length %v and width %v", perimeterTestCase.perimeter, rectangle.length, rectangle.width),
			func(t *testing.T) {
				float64Equals(t, rectangle.getPerimeter(), perimeterTestCase.perimeter)
			})
	}

	for _, areaTestCase := range []AreaTestCase{
		{shape: Rectangle{6.0, 2.0}, area: 12.0},
		{shape: Rectangle{1.2, 2.4}, area: 2.88},
	} {
		rectangle := areaTestCase.shape.(Rectangle)
		t.Run(
			fmt.Sprintf("should return %v when given length %v and width %v", areaTestCase.area, rectangle.length, rectangle.width),
			func(t *testing.T) {
				float64Equals(t, rectangle.getArea(), areaTestCase.area)
			})
	}
}

func TestCircle(t *testing.T) {
	t.Run("implements interface: Shape", func(t *testing.T) {
		var _ Shape = Circle{1.0}
	})
	for _, perimeterTestCase := range []PerimeterTestCase{
		{shape: Circle{1.1}, perimeter: 6.911503837897546},
		{shape: Circle{1.8}, perimeter: 11.309733552923255},
	} {
		circle := perimeterTestCase.shape.(Circle)
		t.Run(
			fmt.Sprintf("should return %v when given a circle of radius %v", perimeterTestCase.perimeter, circle.radius),
			func(t *testing.T) {
				float64Equals(t, circle.getPerimeter(), perimeterTestCase.perimeter)
			})
	}

	for _, areaTestCase := range []AreaTestCase{
		{shape: Circle{6.0}, area: 113.09733552923255},
		{shape: Circle{1.24}, area: 4.830512864159666},
	} {
		circle := areaTestCase.shape.(Circle)
		t.Run(
			fmt.Sprintf("should return %v when given a circle of radius %v", areaTestCase.area, circle.radius),
			func(t *testing.T) {
				float64Equals(t, circle.getArea(), areaTestCase.area)
			})
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("should return 6.911503837897546 when given circle of radius: 1.1", func(t *testing.T) {
		circle := Circle{1.1}
		got := circle.getPerimeter()
		want := 6.911503837897546

		float64Equals(t, got, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("should return 3.8013271108436504 when given circle of radius: 1.1", func(t *testing.T) {
		circle := Circle{1.1}
		got := circle.getArea()
		want := 3.8013271108436504

		float64Equals(t, got, want)
	})
}
