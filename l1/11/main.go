package main

import "fmt"

func main() {
	l1 := []int{1, 2, 3}
	l2 := []int{2, 3, 4}
	m := make(map[int]int) // [число] кол-во повторений
	result := []int{}
	for _, v := range l1 {
		m[v]++
	}
	for _, num := range l2 {
		if m[num] > 0 {
			result = append(result, num)
			m[num]--
		}
	}
	fmt.Println(result)
}
