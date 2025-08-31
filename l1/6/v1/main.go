package main

import (
	"fmt"
	"sync"
	"time"
)

// выход по условию
// горутина закончит работу, когда done = true

func worker(wg *sync.WaitGroup, done *bool) {
	defer wg.Done()
	for !*done {
		fmt.Println("goroutine in process")
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("goroutine end by condition")
}

func main() {
	wg := &sync.WaitGroup{} // для гарантированного завершения работы горутины
	done := false
	wg.Add(1)
	go worker(wg, &done)
	time.Sleep(time.Second * 2)
	done = true
	wg.Wait()
}
