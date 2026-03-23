package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	file, err := os.Open("sample.xml")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var p Person
	err = xml.NewDecoder(file).Decode(&p)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}