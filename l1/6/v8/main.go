package main

import (
	"fmt"
	"sync"
	"time"
)

// закрытие канала

func worker(done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("goroutine end by done chan")
			return
		default:
			fmt.Println("goroutine in process")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	done := make(chan struct{})

	wg.Add(1)
	go worker(done, wg)

	time.Sleep(time.Second * 2) // даем поработать 2 секунды
	close(done)

	wg.Wait()
}
