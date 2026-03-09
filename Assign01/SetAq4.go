package main
import "fmt"
func main(){
	var x int=42;
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:",&x)
	fmt.Printf("Address of x (formatted): %p\n", &x)
}