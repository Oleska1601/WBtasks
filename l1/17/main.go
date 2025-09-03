package main

import "fmt"

func binarySearch(lst []int, el int) int {
	left := 0
	right := len(lst) - 1
	for left <= right {
		idx := (left + right) / 2
		if lst[idx] == el {
			return idx
		} else if lst[idx] < el {
			left = idx + 1
		} else {
			right = idx - 1
		}
	}
	return -1
}

func main() {
	lst := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(lst, 5))  // 5
	fmt.Println(binarySearch(lst, 10)) // -1
}
