package main
import "fmt"
func calculate(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return 
}
func main() {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	s, p := calculate(num1, num2)
	fmt.Println("Sum =", s)
	fmt.Println("Product =", p)
}
