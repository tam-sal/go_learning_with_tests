package astructs

import (
	"fmt"
	"math"
	"testing"
)

var (
	r         Rectangle
	c         Circle
	tolerance float64 = 0.0001
)

func init() {
	r = Rectangle{2.1, 5.2}
	c = Circle{2.6}
}

func TestArea(t *testing.T) {
	checkAreas := func(t testing.TB, got, want float64) {
		t.Helper()
		if math.Abs(got-want) > tolerance {
			t.Errorf("\nGot: %g\nWant: %g", got, want)
		}
	}

	t.Run("Case01: tesing rectangle area", func(t *testing.T) {
		got := r.Area()
		want := 14.6
		checkAreas(t, got, want)
	})
	t.Run("Case02: tesing circle area", func(t *testing.T) {
		got := c.Area()
		want := 21.24
		checkAreas(t, got, want)
	})
}

func ExampleRectangle_Area() {
	r = Rectangle{2.1, 5.2}
	resRectangleArea := r.Area()
	fmt.Printf("%.1f\n", resRectangleArea)
	// Output: 14.6
}

func ExampleCircle_Area() {
	c = Circle{2.6}
	resCircleArea := c.Area()
	fmt.Println(resCircleArea)
	// Output: 21.24
}
