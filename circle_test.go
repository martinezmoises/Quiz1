package main

import "testing"

func TestCircle_Area(t *testing.T) {

	t.Run("circle", func(t *testing.T) {
		circle := Circle{5}
		got := circle.area()
		want := 78.53981633974483

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}

	})
}

func TestCircle_Perimeter(t *testing.T) {

	t.Run("circle", func(t *testing.T) {
		circle := Circle{5}
		got := circle.Perimeter()
		want := 31.415926

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}

	})
}
