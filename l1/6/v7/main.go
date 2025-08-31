package main

import (
	"fmt"
	"sync"
	"time"
)

// канал уведомления
// в канал пишем значение и с благодаря этому горутина прерывается

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
	done <- struct{}{}

	wg.Wait()
}
