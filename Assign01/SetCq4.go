package main
import "fmt"
func main(){
	var num int
	fmt.Println("Enter a Number:")
	fmt.Scan(&num)
	if num >= -9 && num <= 9 {
		fmt.Println("Enter number is a single-digit number.")
	} else {
		fmt.Println("Enter number is NOT a single-digit number.")
	}
}