package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string) {

	time.Sleep(2 * time.Second)
	ch <- "Hello from piyusha!"
}

func main() {

	ch := make(chan string)

	go sendMessage(ch)

	message := <-ch
	fmt.Println(message)
}
