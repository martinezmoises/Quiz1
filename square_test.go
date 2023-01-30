package main

import "testing"

func TestSquare_Area(t *testing.T) {

	result, _ := Square(5.0) //is known as the Blank Identifier, used it to ignore the perimeter return value
	got := result
	want := 25.0

	if got != want {
		t.Errorf("got %g want %g,", got, want)
	}
}

func TestSquare_Perimeter(t *testing.T) {
	_, result := Square(5.0) //used it to ignore the area return value
	got := result
	want := 20.0

	if got != want {
		t.Errorf("got %g want %g,", got, want)
	}
}
