package main

import "fmt"

// Function to double a number (Call by Reference)
func double(n *int) {
    *n = *n * 2
    fmt.Println("Inside double function:", *n) // The value pointed by n is modified
}

func main() {
    num := 7

    fmt.Println("Before double function:", num)
    double(&num) // Passing the address of num to the function

    fmt.Println("After double function:", num) // num in main is modified
}
