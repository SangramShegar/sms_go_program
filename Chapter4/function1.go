package main
import "fmt"
func area (lenght,width int)int{
	Ar:=lenght*width
	return Ar
}
func main(){
	fmt.Printf("Area of rectangle is:%d",area(15,10))
}