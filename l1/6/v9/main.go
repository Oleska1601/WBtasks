package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// закрытие канала

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("goroutine in process")
	time.Sleep(time.Millisecond * 500) //даем поработать 500ms
	fmt.Println("goroutine end by runtime.Goexit")
	runtime.Goexit()
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go worker(wg)

	wg.Wait()
}
