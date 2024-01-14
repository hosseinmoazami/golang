package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ID int) {
	fmt.Printf("Worker %d started\n", ID)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d Done\n", ID)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}

	wg.Wait()
}
