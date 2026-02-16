package main
import "fmt"
func changeValue(x int) {
	x = 100
	fmt.Println("Inside function, x =", x)
}
func main() {
	num := 10
	fmt.Println("Before function call, num =", num)
	changeValue(num)
	fmt.Println("After function call, num =", num)
}
