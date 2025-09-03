package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	m := make(map[rune]struct{})
	str := strings.ToLower(s)
	for _, v := range str {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aAbcd"
	fmt.Println(isUnique(s1))
	fmt.Println(isUnique(s2))
	fmt.Println(isUnique(s3))
}
