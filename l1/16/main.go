package main

import (
	"fmt"
	"math/rand"
)

func quickSort(lst []int) []int {
	if len(lst) <= 1 {
		return lst
	}
	pivot := (len(lst) - 1) / 2
	left := []int{}  // содержит элементы < pivot
	right := []int{} // содержит элементы >= pivot
	for i, v := range lst {
		if i != pivot {
			if v < lst[pivot] {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}
	}
	return append(append(quickSort(left), lst[pivot]), quickSort(right)...)
}

func main() {
	lst := make([]int, 0, 10)
	for range 10 {
		lst = append(lst, rand.Intn(10))
	}
	res := quickSort(lst)
	fmt.Println(lst, res)
}
