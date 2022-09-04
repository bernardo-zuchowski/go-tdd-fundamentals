package structs

import "testing"

var rectangle = Rectangle{10.0, 10.0}
var circle = Circle{10.0}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasPerimeter: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
		{name: "RightTriangle", shape: RightRectangle{Base: 10, Height: 5}, hasPerimeter: 26.18033988749895},
	}

	for _, test := range perimeterTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Perimeter()
			if got != test.hasPerimeter {
				t.Errorf("%#v got %g hasPerimeter %g", test.shape, got, test.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "RightTriangle", shape: RightRectangle{Base: 10, Height: 5}, hasArea: 25},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("%#v got %g hasArea %g", test.shape, got, test.hasArea)
			}
		})
	}
}
