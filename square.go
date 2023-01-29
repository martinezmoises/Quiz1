package main

import (
	"fmt"
	"log"
	"math"
)

//Write a Go function named 'square' that accepts a parameter of type float64 that represents the side of the square.
//The 'square' function should return the area and the perimeter of the square when called.

func Square(a float64) (float64, float64) {

	area := math.Pow(a, 2)
	perimeter := 4 * a

	return area, perimeter
}

func main() {

	log.Printf("The area and perimeter of the square is: ")
	fmt.Println(Square(5))
}
