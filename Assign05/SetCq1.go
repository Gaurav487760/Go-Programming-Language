package main

import "fmt"

type Event struct {
	Name string
	Age  int
}

func handler(e Event) {
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
}

func main() {

	event := Event{"Rahul", 21}
	handler(event)
}
