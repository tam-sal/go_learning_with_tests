package structs

func Perimeter(rectangle Rectangle) float64 {
	if rectangle.Width <= 0 || rectangle.Height <= 0 {
		return -1.0
	}
	return (rectangle.Width + rectangle.Height) * 2
}

func Area(rectangle Rectangle) float64 {
	if rectangle.Width <= 0 || rectangle.Height <= 0 {
		return -1.0
	}
	return rectangle.Width * rectangle.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}
