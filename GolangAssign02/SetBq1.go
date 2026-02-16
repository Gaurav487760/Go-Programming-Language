package main
import "fmt"
func swap(a *int, b *int) {
	*a, *b = *b, *a
}
func main() {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Println("Before Swapping:")
	fmt.Println("num1 =", num1, "num2 =", num2)
	// Passing address of variables
	swap(&num1, &num2)
	fmt.Println("After Swapping:")
	fmt.Println("num1 =", num1, "num2 =", num2)
}
