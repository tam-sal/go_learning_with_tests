package cstructs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return math.Round((r.Height+r.Width)*2*100) / 100
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Trunc(c.Radius*c.Radius*math.Pi*100) / 100
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return math.Round(t.Base*t.Height*0.5*100) / 100
}
