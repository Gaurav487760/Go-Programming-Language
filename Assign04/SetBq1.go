package main

import "fmt"

type Student struct {
	name string
	roll int
}

func (s *Student) show() {
	fmt.Println("Student Name:", s.name)
	fmt.Println("Roll Number:", s.roll)
}

func main() {
	s1 := Student{"Rahul", 101}
	s1.show()
}
