package main

import "fmt"

type Number struct {
    value int
}

func (n Number) AddOne() {
    n.value++
    fmt.Println("Inside AddOne (value receiver):", n.value)
}

func (n *Number) AddOnePointer() {
    n.value++
    fmt.Println("Inside AddOnePointer (pointer receiver):", n.value)
}

func main() {
    num := Number{value: 10}

    num.AddOne()     
    fmt.Println("After AddOne:", num.value) 

    num.AddOnePointer()
    fmt.Println("After AddOnePointer:", num.value) 
}
