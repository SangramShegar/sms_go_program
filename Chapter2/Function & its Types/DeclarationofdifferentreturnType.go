package main
import "fmt"
type Person struct{
	Name string
	Age int
}
func createPerson(name string,age int) Person{
	return Person{Name:name,Age:age}
}
func main(){
	p:= createPerson("John",30)
	fmt.Println("Person:",p.Name,p.Age)
}