package main

import (
	"fmt"
	"math"
)

//define interfaces
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, hieght float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.width * r.hieght
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("hello this is go")
}
