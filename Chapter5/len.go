package main

import "fmt"

func main() {
	ch := make(chan int, 5) 
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("Current length of channel:", len(ch)) 
}