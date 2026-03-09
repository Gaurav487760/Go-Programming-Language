package main
import "fmt"
func main(){
	var a, b int
	fmt.Println("Enter a first Number")
	fmt.Scan(&a)
	fmt.Println("Enter a Second Number")
	fmt.Scan(&b)

	fmt.Println("Before Swapping")
	fmt.Println("a=",a)
	fmt.Println("b=",b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After Swapping")
	fmt.Println("a=",a)
	fmt.Println("b=",b)

}