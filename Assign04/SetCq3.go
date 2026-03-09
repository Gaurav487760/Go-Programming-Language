package main

import "fmt"

type Printer interface {
	print()
}

type Scanner interface {
	scan()
}

type Machine interface {
	Printer
	Scanner
}

type Device struct{}

func (d Device) print() {
	fmt.Println("Printing document...")
}

func (d Device) scan() {
	fmt.Println("Scanning document...")
}

func main() {

	var m Machine
	m = Device{}

	m.print()
	m.scan()
}
