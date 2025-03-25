package main

import "fmt"


func calculateCircleArea(radius float64) float64 {
    return 3.14159 * radius * radius
}

func main() {
   
    circleArea := calculateCircleArea(7) 
    
    fmt.Println("Area of the circle:", circleArea)
}
