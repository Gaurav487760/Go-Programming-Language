package main
import "fmt"
func main() {
	slice := []int{
		10,
		20,
		30,
		40,
		50,
	}
	fmt.Println("Slice elements are:")
	for _, v := range slice {
		fmt.Println(v)
	}
}