package main

//отмена с каналом

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func worker(stopCh chan struct{}, wg *sync.WaitGroup, i int, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		// работа остановится при непосредственном обнаружении закрытие канала stopCh, сигнализирующем об остановке,
		// или при закрытии канала ch в главной горутине main
		case <-stopCh:
			fmt.Printf("stop worker %d due to stopCh\n", i)
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
	stopCh := make(chan struct{})

	if len(os.Args) < 2 {
		log.Fatalln("no workers were provided")
	}

	workers, err := strconv.Atoi(os.Args[1])
	if err != nil || workers <= 0 {
		log.Fatalln("incorrect workers")
	}

	ch := make(chan int)

	wg := &sync.WaitGroup{}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT)

	for i := range workers {
		wg.Add(1)
		go worker(stopCh, wg, i, ch)
	}

	go func() {
		defer close(ch)
		for {
			select {
			case <-stopCh: // вызов сигнала отмены -> закрытие канала stopCh, сигнализирующем об остановке -> закрытие канала для записи ch
				return
			case ch <- rand.Intn(1000): //постоянная запись данных в канал
			}

		}
	}()
	<-exit
	close(stopCh)
	wg.Wait()

}
