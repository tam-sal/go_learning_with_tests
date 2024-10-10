package bstructs

import (
	"math"
	"testing"
)

var (
	r         Rectangle
	c         Circle
	tolerance float64 = 0.0001
)

type Shape interface {
	Area() float64
}

func init() {
	r = Rectangle{2.1, 5.2}
	c = Circle{2.6}
}
func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, s Shape, want float64) {
		t.Helper()
		got := s.Area()
		if math.Abs(want-got) > tolerance {
			t.Errorf("\nGot: %g \nWant: %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		checkArea(t, r, 14.6)
	})
	t.Run("circles", func(t *testing.T) {
		checkArea(t, c, 21.24)
	})
}
