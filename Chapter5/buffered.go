package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string) {
	
	time.Sleep(1 * time.Second)
	ch <- "Hello from piyusha!"
}

func main() {
	
	ch := make(chan string, 2)

	ch <- "Message 1"

	go sendMessage(ch)
	
	message1 := <-ch
	message2 := <-ch

	fmt.Println(message1) 
	fmt.Println(message2)
}
