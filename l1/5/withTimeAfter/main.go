package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func writer(stop chan struct{}) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case <-stop:
				return
			case ch <- rand.Intn(1000):
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
	timeout := time.After(time.Duration(t) * time.Second)
	for {
		select {
		case <-timeout:
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
