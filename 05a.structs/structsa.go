package astructs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return (r.Height + r.Width) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	area := math.Round(c.Radius*c.Radius*math.Pi*100) / 100
	return area
}
