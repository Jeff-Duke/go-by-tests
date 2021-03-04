package geometry

import "math"

// Rectangle is a basic shape with 4 sides
type Rectangle struct {
	Length float64
	Width  float64
}

// Circle is a basic shape
type Circle struct {
	Radius float64
}

// Perimeter calculates a perimeter
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Length + rect.Width)
}

// Returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// // Area returns the area of a shape
// func Area(rect Rectangle) float64 {
// 	return rect.Length * rect.Width
// }
