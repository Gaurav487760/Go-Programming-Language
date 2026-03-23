package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
}