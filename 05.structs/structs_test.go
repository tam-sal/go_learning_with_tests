package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	h1 := 0.5
	w1 := 0.6
	h2 := 1.2
	w2 := 3.5
	h3 := 0.0
	w3 := 4.0
	h4 := 5.5
	w4 := -0.5
	h5 := 5.0
	w5 := 2.0
	r1 := Rectangle{w1, h1}
	r2 := Rectangle{w2, h2}
	r3 := Rectangle{w3, h3}
	r4 := Rectangle{w4, h4}
	r5 := Rectangle{w5, h5}

	checkPerimeters := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %g\nWant: %g", got, want)
		}
	}

	t.Run("Case01: perimeter of 0.x cases", func(t *testing.T) {
		got := Perimeter(r1)
		want := 2.20
		checkPerimeters(t, got, want)
	})
	t.Run("Case02: perimeter of x.x cases", func(t *testing.T) {
		got := Perimeter(r2)
		want := 9.40
		checkPerimeters(t, got, want)
	})
	t.Run("Case03: perimeter of h=0 cases", func(t *testing.T) {
		got := Perimeter(r3)
		want := -1.0
		checkPerimeters(t, got, want)
	})
	t.Run("Case04: perimeter of w=-1 cases", func(t *testing.T) {
		got := Perimeter(r4)
		want := -1.0
		checkPerimeters(t, got, want)
	})
	t.Run("Case05: perimeter of whole numbers h and w cases", func(t *testing.T) {
		got := Perimeter(r5)
		want := 14.00
		checkPerimeters(t, got, want)
	})

}

func ExamplePerimeter() {
	h1 := 0.5
	w1 := 0.6
	h2 := 1.2
	w2 := 3.5
	h3 := 0.0
	w3 := 4.0
	h4 := 5.5
	w4 := -0.5
	h5 := 5.0
	w5 := 2.0
	rec1 := Rectangle{w1, h1}
	rec2 := Rectangle{w2, h2}
	rec3 := Rectangle{w3, h3}
	rec4 := Rectangle{w4, h4}
	rec5 := Rectangle{w5, h5}
	res1 := Perimeter(rec1)
	res2 := Perimeter(rec2)
	res3 := Perimeter(rec3)
	res4 := Perimeter(rec4)
	res5 := Perimeter(rec5)

	fmt.Println(res1, res2, res3, res4, res5)
	// Output: 2.2 9.4 -1 -1 14
}

func TestArea(t *testing.T) {
	h1 := 0.5
	w1 := 0.6
	h2 := 1.2
	w2 := 3.5
	h3 := 0.0
	w3 := 4.0
	h4 := 5.5
	w4 := -0.5
	h5 := 5.0
	w5 := 2.0

	rec1 := Rectangle{w1, h1}
	rec2 := Rectangle{w2, h2}
	rec3 := Rectangle{w3, h3}
	rec4 := Rectangle{w4, h4}
	rec5 := Rectangle{w5, h5}

	checkAreas := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %g\nWant: %g", got, want)
		}
	}

	t.Run("Case01: Checking fractions", func(t *testing.T) {
		got := Area(rec1)
		want := 0.3
		checkAreas(t, got, want)

	})
	t.Run("Case02: Checking floats", func(t *testing.T) {
		got := Area(rec2)
		want := 4.2
		checkAreas(t, got, want)

	})
	t.Run("Case03: Checking zeros", func(t *testing.T) {
		got := Area(rec3)
		want := -1.0
		checkAreas(t, got, want)

	})
	t.Run("Case04: Checking negative fractions", func(t *testing.T) {
		got := Area(rec4)
		want := -1.0
		checkAreas(t, got, want)

	})
	t.Run("Case05: Checking whole nums", func(t *testing.T) {
		got := Area(rec5)
		want := 10.0
		checkAreas(t, got, want)

	})
}

func ExampleArea() {
	h1 := 0.5
	w1 := 0.6
	h2 := 1.2
	w2 := 3.5
	h3 := 0.0
	w3 := 4.0
	h4 := 5.5
	w4 := -0.5
	h5 := 5.0
	w5 := 2.0

	rec1 := Rectangle{w1, h1}
	rec2 := Rectangle{w2, h2}
	rec3 := Rectangle{w3, h3}
	rec4 := Rectangle{w4, h4}
	rec5 := Rectangle{w5, h5}
	res1 := Area(rec1)
	res2 := Area(rec2)
	res3 := Area(rec3)
	res4 := Area(rec4)
	res5 := Area(rec5)
	fmt.Println(res1, res2, res3, res4, res5)
	// Output: 0.3 4.2 -1 -1 10
}
