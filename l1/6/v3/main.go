package main

import (
	"fmt"
	"sync"
	"time"
)

// timer
// горутина закончит работу, когда истечет timer (в канал timer.C отправится текущее время)

func worker(timer *time.Timer, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-timer.C:
			fmt.Println("goroutine end by timer")
			return
		default:
			fmt.Println("goroutine in process")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	timer := time.NewTimer(time.Second * 2)
	wg.Add(1)
	go worker(timer, wg)
	wg.Wait()
}
