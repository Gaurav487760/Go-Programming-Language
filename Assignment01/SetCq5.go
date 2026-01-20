package main
import (
	"fmt"
	"strings"
)
func main() {
	var s1, s2 string

	fmt.Print("Enter first string: ")
	fmt.Scan(&s1)

	fmt.Print("Enter second string: ")
	fmt.Scan(&s2)

	if strings.Contains(s2, s1) {
		fmt.Println("Substring found")
	} else {
		fmt.Println("Substring not found")
	}
}