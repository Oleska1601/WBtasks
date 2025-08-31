package main

import (
	"fmt"
	"sync"
	"time"
)

// time.After
// горутина закончит работу, когда истечет time.After (в канал, созданный с помощью time.After, отправится текущее время)

func worker(timeout <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-timeout:
			fmt.Println("goroutine end by time.After")
			return
		default:
			fmt.Println("goroutine in process")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	timeout := time.After(time.Second * 2)
	wg.Add(1)
	go worker(timeout, wg)
	wg.Wait()
}
