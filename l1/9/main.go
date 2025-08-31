package main

import "fmt"

func generator(lst []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, v := range lst {
			ch <- v
		}
	}()

	return ch
}

func doubler(inputCh <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range inputCh {
			ch <- v * 2
		}
	}()
	return ch
}

func main() {
	lst := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputCh := generator(lst)
	doubleCh := doubler(inputCh)
	for v := range doubleCh {
		fmt.Println(v)
	}

}
