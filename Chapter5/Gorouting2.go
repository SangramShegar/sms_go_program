package main
import "fmt"
func main(){
go message()
	fmt.Print("hii!from main go routine")
}
func message(){
fmt.Print("hii!from message go routine")
}