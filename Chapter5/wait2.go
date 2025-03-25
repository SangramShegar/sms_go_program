package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() 
	fmt.Printf("Worker %d: Started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d: Finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	wg.Wait() 
	fmt.Println("All workers completed.")
}