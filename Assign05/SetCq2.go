package main

import (
	"fmt"
	"time"
)

func display(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	go display("Goroutine 1")
	go display("Goroutine 2")

	time.Sleep(time.Second * 3)
}
