package main
import "fmt"
func main(){
	var rows int
	fmt.Println("Enter number of Rows:")
	fmt.Scan(&rows)
	for i := 0; i < rows; i++ {
		num := 1
		for space := 1; space <= rows-i; space++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(num, " ")
			num = num * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}