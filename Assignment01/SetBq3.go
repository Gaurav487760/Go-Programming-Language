package main
import "fmt"
func main(){
	var n int
	var first, second int = 0, 1

	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)

	fmt.Println("Fibonacci Series:")
	for i := 1; i <= n; i++ {
		fmt.Print(first, " ")
		next := first + second
		first = second
		second = next
	}
}