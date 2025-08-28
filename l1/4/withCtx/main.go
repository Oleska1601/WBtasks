package main

//отмена с контекстом

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
)

func worker(ctx context.Context, wg *sync.WaitGroup, i int, ch <-chan int64) {
	defer wg.Done()
	for {
		select {
		// работа остановится при непосредственном обнаружении отмены контекста или при закрытии канала ch в главной горутине main
		case <-ctx.Done():
			fmt.Printf("stop worker %d due to context\n", i)
			return
		case v, ok := <-ch:
			if !ok {
				fmt.Printf("stop worker %d due to close channel\n", i)
				return
			}
			fmt.Printf("worker %d process value %d\n", i, v)
		}
	}

}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("no workers were provided")
	}

	workers, err := strconv.Atoi(os.Args[1])
	if err != nil || workers <= 0 {
		log.Fatalln("incorrect workers")
	}

	ch := make(chan int64)
	var v atomic.Int64

	wg := &sync.WaitGroup{}
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()
	for i := range workers {
		wg.Add(1)
		go worker(ctx, wg, i, ch)
	}

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done(): // вызов сигнала отмены -> отмена контекста -> закрытие канала для записи ch
				return
			case ch <- v.Add(1): //постоянная запись данных в канал
			}

		}
	}()

	wg.Wait()

}
