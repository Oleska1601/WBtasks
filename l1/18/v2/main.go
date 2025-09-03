package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count atomic.Int64
}

func (c *Counter) Inc() {
	c.count.Add(1)
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
	fmt.Println(counter.count.Load())
}
