package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle has a method called Perimeter that returns a float64 so it satisfies the Shape interface
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Circle has a method called Perimeter that returns a float64 so it satisfies the Shape interface
func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

// Circle has a method called Area that returns a float64 so it satisfies the Shape interface
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type RightRectangle struct {
	Base   float64
	Height float64
}

// RightTriangle has a method called Perimeter that returns a float64 so it satisfies the Shape interface
func (t RightRectangle) Perimeter() float64 {
	hypotenuse := math.Hypot(t.Base, t.Height)

	return t.Base + t.Height + hypotenuse
}

// RightTriangle has a method called Area that returns a float64 so it satisfies the Shape interface
func (t RightRectangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
