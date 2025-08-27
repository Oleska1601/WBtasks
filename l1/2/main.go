package main

import (
	"fmt"
	"sync"
)

func main() {
	sl := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	for _, v := range sl {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Printf("square of %d is %d\n", v, v*v)
		}(v)
	}
	wg.Wait()
}
