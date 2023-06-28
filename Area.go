package main

import (
	"fmt"
	"math"
)

// shape i an interface that defines a method for calculating the area
type Shape interface {
	Area() float64
}

// Circle is a struct representing a circle
type Circle struct {
	radius float64
}

// Area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius

}

// Rectangle is a struct representing a rectangle
type Rectangle struct {
	width  float64
	height float64
}

// Area calculates the area of the  rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// printArea calculatess and prints the area of the given shape
func PrintArea(s Shape) {
	fmt.Printf("Area:%0.2f\n", s.Area())
}
func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}
	PrintArea(circle)
	PrintArea(rectangle)

}