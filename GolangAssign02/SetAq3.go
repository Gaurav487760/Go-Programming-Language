package main
import "fmt"
func isPalindrome(n int) bool {
	original := n
	reverse := 0
	for n > 0 {
		remainder := n % 10
		reverse = reverse*10 + remainder
		n = n / 10
	}
	return original == reverse
}
func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	if isPalindrome(number) {
		fmt.Println("The number is a Palindrome")
	} else {
		fmt.Println("The number is NOT a Palindrome")
	}
}
