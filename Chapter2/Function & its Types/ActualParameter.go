package main

import "fmt"


func add(x int, y int) int {
    return x + y
}

func main() {
    a := 5   
    b := 10  

    result := add(a, b)  
    fmt.Println("Sum:", result)
}
