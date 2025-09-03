package main

import "fmt"

func rotate(lst []rune, l, r int) {
	for l < r {
		lst[l], lst[r] = lst[r], lst[l]
		l++
		r--
	}
}

func rotateSentence(s string) string {
	lst := []rune(s)
	rotate(lst, 0, len(s)-1) // переворот всей строки
	start := 0
	var i int
	for i = 1; i <= len(lst); i++ {
		if i == len(lst) || lst[i] == ' ' {
			rotate(lst, start, i-1) // переворот каждого слова
			start = i + 1
		}

	}
	return string(lst)
}

func main() {
	s := "sun dog snow "
	fmt.Println(rotateSentence(s))
}
