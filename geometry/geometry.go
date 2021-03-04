package geometry

import "math"

// Perimeter calculates a perimeter
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Length + rect.Width)
}

// Shape is an interface for more specific shapes
type Shape interface {
	Area() float64
}

// Rectangle is a basic shape with 4 sides
type Rectangle struct {
	Length float64
	Width  float64
}

// Circle is a basic shape
type Circle struct {
	Radius float64
}

// Triangle is a basic shape
type Triangle struct {
	Base, Height float64
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area returns the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
