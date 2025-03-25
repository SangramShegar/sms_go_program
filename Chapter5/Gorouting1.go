package main
import "fmt"
func main(){
	message()
	fmt.Print("hii!from main go routine")
}
func message(){
fmt.Print("hii!from message go routine")
}