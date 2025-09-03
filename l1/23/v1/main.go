package main

import (
	"fmt"
)

func remove(sl []int, i int) []int {
	copy(sl[i:], sl[i+1:])
	return sl[:len(sl)-1]
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7}
	sl = remove(sl, 4)
	fmt.Println(sl)
}
