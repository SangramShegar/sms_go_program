package main

import "fmt"

func main() {
	ch := make(chan string, 4) 
	fmt.Println("Channel capacity:", cap(ch)) 
}