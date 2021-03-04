package geometry

type Rectangle struct {
	Length float64
	Width  float64
}

// Perimeter calculates a perimeter
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Length + rect.Width)
}

// Area returns the area of a shape
func Area(rect Rectangle) float64 {
	return rect.Length * rect.Width
}
