package main

import "fmt"

func findSet(lst []string) []string {
	result := []string{}
	m := make(map[string]struct{})
	for _, v := range lst {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = struct{}{}
		}
	}

	return result

}

func main() {
	lst := []string{"cat", "cat", "dog", "cat", "tree"}
	result := findSet(lst)
	fmt.Println(result)

}
