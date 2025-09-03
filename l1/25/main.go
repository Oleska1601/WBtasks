package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	now := time.Now()
	sleep(time.Second * 3)
	fmt.Println("time duration: ", time.Since(now))
}
