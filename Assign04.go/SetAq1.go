package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{length: 10, breadth: 4}

	fmt.Println("Circle Area:", c.area())
	fmt.Println("Circle Perimeter:", c.perimeter())

	fmt.Println("Rectangle Area:", r.area())
	fmt.Println("Rectangle Perimeter:", r.perimeter())
}
