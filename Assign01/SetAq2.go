package main
import "fmt"

func main(){
	var num int
	fmt.Println("Enter a Number:")
	fmt.Scan(&num)
	if num%2==0{
		fmt.Println("The Number is Even")
		}else{
			fmt.Println("The Number is odd")
		}
	}