package main

import "fmt"

type Author struct {
	name  string
	books int
}

func (a Author) Show() {
	fmt.Printf("Author Name: %s, Books Written: %d\n", a.name, a.books)
}

func main() {
	author := Author{name: "PIYUSHA GARAD", books: 3}
	author.Show()
}
