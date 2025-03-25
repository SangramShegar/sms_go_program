package main

import "fmt"

func sender(ch chan string) {
	fmt.Println("Sender: Sending message...")
	ch <- "Hello, Go Channels!"
	fmt.Println("Sender: Message sent!")
}

func receiver(ch chan string) {
	fmt.Println("Receiver: Waiting for message...")
	msg := <-ch
	fmt.Println("Receiver: Received message:", msg)
}

func main() {
	ch := make(chan string) 
	
	go sender(ch)
	go receiver(ch)

	fmt.Scanln()
}