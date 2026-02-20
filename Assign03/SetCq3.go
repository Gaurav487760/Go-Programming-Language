package main
import "fmt"

func main() {

	s := []int{10, 20, 30}
	fmt.Println("Original Slice:", s)

	s = append(s, 40)
	fmt.Println("After Append:", s)

	index := 1
	s = append(s[:index], s[index+1:]...)
	fmt.Println("After Remove:", s)

	newSlice := make([]int, len(s))
	copy(newSlice, s)
	fmt.Println("Copied Slice:", newSlice)
}