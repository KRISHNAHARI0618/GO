package main

import (
	"fmt"
	"math"
)

// Declaring the Structs
type square struct {
	length float64
}

type circle struct {
	radius float64
}

// Declaring the methods

// Square Methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// Circle Methods

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// Declaring Interface with two methods are including

type shaping interface {
	area() float64
	circumf() float64
}

// printing the shaping function

func shapingFunc(s shaping) {
	fmt.Printf("Area Of %T is : %.02f \n", s, s.area())
	fmt.Printf("Circumf of %T is: %0.02f \n", s, s.circumf())

}

func main(s shaping) {
	fmt.Println("Peddireddy Hari Vardhan Redyy....")

	shapes := []shaping{
		square{length: 4},
		circle{radius: 7},
		circle{radius: 8},
		square{length: 10},
	}

	fmt.Println(shapes)
}
