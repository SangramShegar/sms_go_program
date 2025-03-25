package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from Channel 1"
	}()

	
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from Channel 2"
	}()


	for i := 0; i < 3; i++ { 
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		default:
			fmt.Println("No message received yet... checking again")
			time.Sleep(500 * time.Millisecond)
		}
	}

	fmt.Println("Program finished")
}