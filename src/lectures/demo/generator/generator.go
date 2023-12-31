package main

import (
	"fmt"
	"math/rand"
)

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max - min + 1)
		}
	}()

	return out
}

func generateRandIntN(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()
	return out
}

func main() {
	randInt := generateRandInt(1, 100)

	fmt.Println("generateRandInt infinite")

	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println("generateRandInt using infinite count")
	randIntRange := generateRandIntN(2, 1, 10)

	for i := range randIntRange {
		fmt.Println(i)
	}

	fmt.Println(<-randIntRange)
	fmt.Println(<-randIntRange)
	fmt.Println(<-randIntRange)
	fmt.Println(<-randIntRange)
	fmt.Println(<-randIntRange)
	fmt.Println(<-randIntRange)

	fmt.Println("generateRandInt using infinite loop")

	randIntn := generateRandIntN(3, 1, 10)
	for {
		n, open := <-randIntn
		if !open {
			break
		}
		fmt.Println(n)
	}
}
