package main

import "fmt"

type Person interface{}

func displayDetails(p Person) {
	if val, ok := p.(string); ok {
		fmt.Println("Person Name:", val)
	} else {
		fmt.Println("Not a string type")
	}
}
func main() {
	var p Person = "Piyusha Garad"
	displayDetails(p)
}
