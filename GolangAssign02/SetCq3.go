package main
import "fmt"
func calculate(a int, b int) (int, int, int) {
	sum := a + b
	diff := a - b
	product := a * b
	return sum, diff, product
}
func main() {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	s, d, p := calculate(num1, num2)
	fmt.Println("Sum =", s)
	fmt.Println("Difference =", d)
	fmt.Println("Product =", p)
}
