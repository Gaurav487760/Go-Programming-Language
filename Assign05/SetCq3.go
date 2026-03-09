package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}

func cube(n int, ch chan int) {
	ch <- n * n * n
}

func main() {

	num := 3

	sq := make(chan int)
	cb := make(chan int)

	go square(num, sq)
	go cube(num, cb)

	squareVal := <-sq
	cubeVal := <-cb

	sum := squareVal + cubeVal

	fmt.Println("Square:", squareVal)
	fmt.Println("Cube:", cubeVal)
	fmt.Println("Sum:", sum)
}
