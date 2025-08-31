package main

import "fmt"

func main() {
	lst := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float32)
	for _, v := range lst {
		key := int(v/10) * 10
		m[key] = append(m[key], v)
	}
	fmt.Println(m)
}
