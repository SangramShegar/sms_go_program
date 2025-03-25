package main

import "fmt"

type Array struct {
	elements [5]int
}
func (a *Array) Copy(b [5]int) {
	for i := 0; i < len(b); i++ {
		a.elements[i] = b[i]
	}
}
func main() {
	source := [5]int{1, 2, 3, 4, 5}
	var destination Array

	destination.Copy(source)

	fmt.Println("Copied Array:", destination.elements)
}
