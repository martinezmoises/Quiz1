package main

import "testing"

func TestSquare_Area(t *testing.T) {
	result, _ := Square(5.0)
	got := result
	want := 25.0

	if got != want {
		t.Errorf("got %g want %g,", want, got)
	}
}

func TestSquare_Perimeter(t *testing.T) {
	_, result := Square(5.0)
	got := result
	want := 20.0

	if got != want {
		t.Errorf("got %g want %g,", want, got)
	}
}
