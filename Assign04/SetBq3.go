package main

import "fmt"

type Array struct{}

func (a Array) copyArray(src [5]int) [5]int {
	var dest [5]int

	for i := 0; i < len(src); i++ {
		dest[i] = src[i]
	}
	return dest
}

func main() {
	src := [5]int{1, 2, 3, 4, 5}
	var a Array

	dest := a.copyArray(src)

	fmt.Println("Source Array:", src)
	fmt.Println("Copied Array:", dest)
}
