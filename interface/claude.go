// Create a Shape interface with an Area() float64 method. Then:

// Create a Circle struct with a radius field
// Create a Rectangle struct with width and height fields
// Implement the Area() method for both
// Write a function printAreas that takes a slice of Shape and prints each area

// Formulas:

// Circle area = 3.14 * radius * radius
// Rectangle area = width * height

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct{
	radius float64
}

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func PrintArea(p []Shape) {
	for _, charset:=range p {
		fmt.Println(charset.Area())
	}
}

func mainn() {

	c1 := Circle{radius: 32}
	r1 := Rectangle{height: 4, width: 3}
	shapes := []Shape{r1, c1}
	 PrintArea(shapes)
}



