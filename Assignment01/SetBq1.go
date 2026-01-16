package main
import "fmt"
func main(){
	var num int

	fmt.Println("Enter a Number")
	fmt.Scan(&num)

	fmt.Println("Table of", num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}
}