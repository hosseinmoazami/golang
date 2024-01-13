package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 10)
	queue <- "one"
	queue <- "two"
	close(queue)

	for item := range queue {
		fmt.Println(item)
	}
}
