package main
import "fmt"
func calculate(x int ,y int)(int,int){
	sum := x+y
	diff :=x-y
	return sum,diff
}
func main(){
	a:=10
	b:=5
	sum,diff :=calculate(a,b)
	fmt.Println("sum",sum)
	fmt.Println("difference",diff)
}