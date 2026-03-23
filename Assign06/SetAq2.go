package main

import (
	"fmt"
	"calculator"
)

func main() {

	var a, b, choice int

	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)

	fmt.Println("1.Add 2.Sub 3.Mul 4.Div")
	fmt.Println("Enter choice:")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		fmt.Println("Result:", calculator.Add(a, b))

	case 2:
		fmt.Println("Result:", calculator.Sub(a, b))

	case 3:
		fmt.Println("Result:", calculator.Mul(a, b))

	case 4:
		fmt.Println("Result:", calculator.Div(a, b))

	default:
		fmt.Println("Invalid choice")
	}
}