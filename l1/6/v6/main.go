package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context.WithCancel
// горутина закончит работу по причине закрытия канала ctx.Done() при вызове функции cancel

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine end by context")
			return
		default:
			fmt.Println("goroutine in process")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx, wg)

	time.Sleep(time.Second * 2) // даем поработать 2 секунды
	cancel()

	wg.Wait()
}
