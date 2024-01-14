package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count map[string]int
}

func (c *Counter) Inc(index string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count[index]++
}

func main() {
	c := Counter{
		count: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.Inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)
	go doIncrement("a", 1000)

	wg.Wait()
	fmt.Println(c.count)
}
