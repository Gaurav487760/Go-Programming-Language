package main
import "fmt"
func main(){
var a, b int
	var choice int

	fmt.Println("Enter first number: ")
	fmt.Scan(&a)
	fmt.Println("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	fmt.Println("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Addition =", a+b)
	case 2:
		fmt.Println("Subtraction =", a-b)
	case 3:
		fmt.Println("Multiplication =", a*b)
	case 4:
		if b != 0 {
			fmt.Println("Division =", a/b)
		} else {
			fmt.Println("Division by zero not allowed")
		}
	default:
		fmt.Println("Invalid choice")
	}
}
