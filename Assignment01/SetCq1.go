package main
import "fmt"
func main(){
	var str1, str2 string

	fmt.Println("Enter first string: ")
	fmt.Scan(&str1)
	fmt.Println("Enter second string: ")
	fmt.Scan(&str2)

	p1 := &str1
	p2 := &str2

	result := *p1 + *p2

	fmt.Println("Concatenated String:", result)
}