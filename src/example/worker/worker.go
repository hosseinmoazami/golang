package main

import (
	"fmt"
	"time"
)

func worker(ID int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", ID, "Stared job", job)
		time.Sleep(time.Second)
		fmt.Println("Worker", ID, "Finish job", job)
		results <- 2 * job
	}
}

func main() {
	const NumJobs = 5
	jobs := make(chan int, NumJobs)
	results := make(chan int, NumJobs)

	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= NumJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 0; r < NumJobs; r++ {
		<-results
	}
}
