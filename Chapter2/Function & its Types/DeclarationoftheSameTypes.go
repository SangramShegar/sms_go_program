package main
import "fmt"
func multiply(x,y,z int)int{
	return x*y*z
}
func main(){
	result:=multiply(2,3,4)
	fmt.Println("product",result)
}