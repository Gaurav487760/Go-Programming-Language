package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("sample.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	content := "This is appended text.\n"

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	fmt.Println("Content appended successfully!")
}