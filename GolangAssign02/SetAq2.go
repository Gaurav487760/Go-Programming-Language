package main
import "fmt"
func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumOfDigits(n/10)
}
func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	result := sumOfDigits(number)
	fmt.Println("Sum of digits =", result)
}
