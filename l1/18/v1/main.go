package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int64
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func main() {
	counter := &Counter{}
	wg := &sync.WaitGroup{}

	for range 10000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()

	}
	wg.Wait()
	fmt.Println(counter.count)
}
