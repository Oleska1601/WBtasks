package main

import "fmt"

func rotate(s string) string {
	lst := []rune(s)
	l := 0
	r := len(lst) - 1
	for l < r {
		lst[l], lst[r] = lst[r], lst[l]
		l++
		r--
	}

	return string(lst)
}

func main() {
	fmt.Println(rotate("hello123"))
	fmt.Println(rotate("привет123.."))
}
