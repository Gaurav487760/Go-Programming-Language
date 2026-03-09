package main

import "fmt"

func main() {

	var i interface{}

	i = 50

	value, ok := i.(int)

	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Type assertion failed")
	}
}
