package main

import "fmt"

type Author struct {
	name  string
	books int
}

func (a Author) show() {
	fmt.Println("Author Name:", a.name)
	fmt.Println("Number of Books:", a.books)
}

func main() {
	a1 := Author{"R.K. Narayan", 10}
	a1.show()
	a2 := Author{"J.K.Patil", 60}
	a2.show()
}
