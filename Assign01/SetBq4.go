package main
import "fmt"
func main(){
	var x int = 10
	var p *int = &x
	var pp **int = &p
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("\nValue of p (address of x):", p)
	fmt.Println("Address of p:", &p)
	fmt.Println("\nValue of pp (address of p):", pp)
	fmt.Println("Value pointed by pp (value of p):", *pp)
	fmt.Println("Value pointed by *pp (value of x):", **pp)
}