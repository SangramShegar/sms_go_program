package main
import "fmt"
func square(x int)int{
	return x*x
}
func main(){
	num := 4
	result := square(num)
	fmt.Println("square:",result)
}