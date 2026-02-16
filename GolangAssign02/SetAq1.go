package main
import "fmt"
func add(a int, b int) int {
	return a + b
}
func main() {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	result := add(num1, num2)
	fmt.Println("Addition =", result)
}
