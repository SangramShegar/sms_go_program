package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() 
	
	fmt.Printf("Task %d is starting...\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Task %d is completed.\n", id)
}

func main() {
	var wg sync.WaitGroup
	
	numTasks := 5
	wg.Add(numTasks)
	
	for i := 1; i <= numTasks; i++ {
		go task(i, &wg)
	}
	
	wg.Wait()
	
	fmt.Println("All tasks completed!")
}
