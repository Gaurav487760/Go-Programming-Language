package main

import "fmt"

type Numbers struct {
	a int
	b int
}

func (n Numbers) multiply() int {
	return n.a * n.b
}

func main() {
	var n Numbers
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&n.a, &n.b)

	result := n.multiply()
	fmt.Println("Multiplication =", result)
}
