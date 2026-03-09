package main

import (
	"fmt"
	"sort"
)

type Student struct {
	rollno     int
	name       string
	percentage float64
}

type ByPercentage []Student

func (s ByPercentage) Len() int {
	return len(s)
}

func (s ByPercentage) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPercentage) Less(i, j int) bool {
	return s[i].percentage > s[j].percentage
}

func main() {

	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter Roll No, Name, Percentage:")
		fmt.Scan(&students[i].rollno, &students[i].name, &students[i].percentage)
	}

	sort.Sort(ByPercentage(students))

	fmt.Println("\nStudents in Descending Order of Percentage:")

	for _, s := range students {
		fmt.Println(s.rollno, s.name, s.percentage)
	}
}
