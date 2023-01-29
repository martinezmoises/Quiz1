package main

import (
	"log"
	"math"
)

//1. Create a struct type named 'triangle' that has two fields named
//base and height both of type float64.

type Triangle struct {
	Base   float64
	Height float64
}

//2. Create a method on type 'triangle' named 'area' that calculates and
//returns the area of a triangle.

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

//Create a method on type 'triangle' named 'perimeter' that calculates and
//returns the perimeter of a triangle.

// To find perimeter we add all sides of triangle, however we are missing one of the sides which is Hypotenuse. Therefore, we should first find Hypotenuse
// to be ale to properly find the perimeter.

// Formula to find Hypotenuse = sqrt a squared + b squared.

func (t Triangle) Perimeter() float64 {

	hypotenuse := math.Sqrt(math.Pow(t.Height, 2) + math.Pow(t.Base, 2))

	return t.Base + t.Height + hypotenuse

}

func main() {

	//Inside the 'main' function create a variable of type 'triangle' Inside the 'main' function create a variable of type 'triangle'

	//Creating 'module1' which is an object from struct 'Triangle', then we declare the field values to calculate them.
	module1 := Triangle{
		Base:   8.6,
		Height: 4.2,
	}

	//Printing methods
	log.Println("The area of the triangle is: ", module1.Area())
	log.Println("THe perimeter of the triangle is: ", module1.Perimeter())
}
