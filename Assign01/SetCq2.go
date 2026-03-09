package main
import "fmt"
func main(){
	var str1, str2 string
	fmt.Println("Enter first string: ")
	fmt.Scan(&str1)
	fmt.Println("Enter second string: ")
	fmt.Scan(&str2)

	if str1 == str2 {
		fmt.Println("Both strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}
}