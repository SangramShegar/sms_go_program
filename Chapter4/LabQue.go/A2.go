package main

import "fmt"

type Numbers struct {
	a, b int
}

func (n Numbers) Multiply() int {
	return n.a * n.b
}

func main() {
	num := Numbers{a: 5, b: 7}
	fmt.Println("Multiplication:", num.Multiply())
}
