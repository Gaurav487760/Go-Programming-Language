package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Marks int
}

func main() {

	students := []Student{
		{"Rahul", 78},
		{"Amit", 85},
		{"Sneha", 65},
		{"Pooja", 90},
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].Marks < students[j].Marks
	})

	fmt.Println("Students sorted by marks:")
	for _, s := range students {
		fmt.Println(s.Name, s.Marks)
	}
}