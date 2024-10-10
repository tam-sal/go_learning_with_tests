package cstructs

import (
	"testing"
)

var (
	r, r1  Rectangle
	c, c1  Circle
	t1, t2 Triangle
)

type Shape interface {
	Area() float64
}

func init() {
	r = Rectangle{2.1, 5.2}
	r1 = Rectangle{5.2, 3.0}
	c = Circle{2.6}
	c1 = Circle{3.2}
	t1 = Triangle{2.5, 3.6}
	t2 = Triangle{3.8, 1.5}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: r, hasArea: 14.6},
		{name: "Rectangle", shape: r1, hasArea: 16.4},
		{name: "Circle", shape: c, hasArea: 21.23},
		{name: "Circle", shape: c1, hasArea: 32.16},
		{name: "Triangle", shape: t1, hasArea: 4},
		{name: "Triangle", shape: t2, hasArea: 2.85},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("Error: %#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
