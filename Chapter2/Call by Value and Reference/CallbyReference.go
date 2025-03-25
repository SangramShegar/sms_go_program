package main

import "fmt"

// Function to increment a number (Call by Reference)
func incrementByReference(num *int) {
    *num = *num + 1
    fmt.Println("Inside function:", *num)
}

func main() {
    value := 10
    incrementByReference(&value) 
    fmt.Println("Outside function:", value) 
}
