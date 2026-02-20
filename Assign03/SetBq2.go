package main
import (
	"fmt"
	"sort"
)
func main() {
	arr := []int{50, 20, 40, 10, 30}
	sort.Ints(arr)
	fmt.Println("Sorted array:", arr)
}