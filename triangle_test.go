package main

import (
	"testing"
)

func TestTriangle_Area_(t *testing.T) {

	//t.Run enables running “subtests”, one for each table entry. These are shown separately when executing go test -v.
	//Another way you can selectively run tests is by passing arguments into the test runner that can be used by the tests.
	//Function signiture
	t.Run("triangle", func(t *testing.T) {

		triangle := Triangle{9, 4.2}
		got := triangle.Area()
		want := 18.06

		if got != want {
			t.Errorf("got %g want %g ", got, want) //I got this but I wanted that
		}

	})
}

func TestTriangle_Perimeter(t *testing.T) {
	t.Run("triangle", func(t *testing.T) {

		triangle := Triangle{8.6, 4.2}
		got := triangle.Perimeter()
		want := 22.37078889120432

		if got != want {
			t.Errorf("got %g want %g ", got, want) //format
		}

	})
}
