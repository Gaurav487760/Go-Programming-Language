package main 
import "fmt"
func main(){
	p := new(int)

	fmt.Println("Value stored at p:", *p)

	*p = 25
	fmt.Println("Updated value stored at p:", *p)
	fmt.Println("Address stored in p:", p)
}