package main

import (
     "fmt"
	 "rectangle"
)

func main() {

	var l, w float64

	fmt.Println("Enter length and width:")
	fmt.Scan(&l, &w)

	area := rectangle.Area(l, w)

	fmt.Println("Area of rectangle:", area)
} 
  













