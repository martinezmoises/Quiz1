package main

import (
	"log"
	"math"
)

//Create a struct type named 'circle' that has one field named 'radius' of type
//float64.

type Circle struct {
	Radius float64
}

//Create a method on type 'circle' named 'area' that calculates and
//returns the area of a circle.

func (c Circle) area() float64 {

	return math.Pi * math.Pow(c.Radius, 2)

}

//create a method on type 'circle' named 'perimeter' that calculates and
//returns the perimeter of a circle.

func (c Circle) Perimeter() float64 {

	const pi = 3.1415926

	return 2 * pi * c.Radius

}

func main() {

	module2 := Circle{
		Radius: 5,
	}

	log.Println("The area of the circle is: ", module2.area())
	log.Println("The perimeter of the circle is: ", module2.Perimeter())

}
