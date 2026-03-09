package main

import (
	"fmt"
	"math/rand"
	"time"
)

func delay() {
	time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
}

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		delay()
	}
}
