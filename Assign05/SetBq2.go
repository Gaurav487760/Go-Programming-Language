package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			c1 <- "from 1"
		}
	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			c2 <- "from 2"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
