package main
import "fmt"

type Author struct {
	name  string
	books int
}
func (a Author) Details() string {
	return a.name + " has written " + fmt.Sprint(a.books) + " books."
}
func main() {
	author := Author{name: "George Orwell", books: 5}
	fmt.Println(author.Details())
}