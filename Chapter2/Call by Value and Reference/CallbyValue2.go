package main

import "fmt"

// Function to square a number (Call by Value)
func square(n int) {
    n = n * n
    fmt.Println("Inside square function:", n) // Local variable n is modified
}

func main() {
    num := 5

    fmt.Println("Before square function:", num)
    square(num) // Passing value of num

    fmt.Println("After square function:", num) // num in main remains the same
}
