package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func writer(stop chan struct{}) chan int64 {
	ch := make(chan int64)
	go func() {
		defer close(ch)
		var v int64
		for {
			select {
			case <-stop:
				return
			case ch <- v:
				v++
			}
		}
	}()
	return ch
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("no time was provided")
	}
	t, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil || t <= 0 {
		log.Fatalln("incorrect time")
	}

	stopCh := make(chan struct{})
	ch := writer(stopCh)
	timeout := time.NewTimer(time.Duration(t) * time.Second)
	for {
		select {
		case <-timeout.C:
			close(stopCh)
			return
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}

}
