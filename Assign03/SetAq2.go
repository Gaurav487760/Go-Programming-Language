package main
import "fmt"
type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}
func main() {
	var n int
	fmt.Print("Enter number of books: ")
	fmt.Scan(&n)
	books := make([]Book, n)
	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details for Book", i+1)
		fmt.Println("\nEnter ID of a Book")
		fmt.Scan(&books[i].BookID)
		fmt.Println("\nEnter Title of a Book")
		fmt.Scan(&books[i].Title)
		fmt.Println("\nEnter Author of a Book")
		fmt.Scan(&books[i].Author)
		fmt.Println("\nEnter Price of a Book")
		fmt.Scan(&books[i].Price)
	}
	fmt.Println("\nBook Details:")
	for i := 0; i < n; i++ {
		fmt.Println("BookID:", books[i].BookID)
		fmt.Println("Title:", books[i].Title)
		fmt.Println("Author:", books[i].Author)
		fmt.Println("Price:", books[i].Price)
		fmt.Println()
	}
}