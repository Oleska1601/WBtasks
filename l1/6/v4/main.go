package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context.WithTimeout
// горутина закончит работу по причине context.DeadlineExceeded (канал ctx.Done() закроется по истечении времени)

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	wg.Add(1)
	go worker(ctx, wg)
	wg.Wait()
}
