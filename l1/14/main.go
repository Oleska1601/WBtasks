package main

import "fmt"

func recogniseType(input interface{}) {
	switch input.(type) {
	case int:
		fmt.Println("input is type int")
	case string:
		fmt.Println("input is type string")
	case bool:
		fmt.Println("input is type bool")
	case chan int:
		fmt.Println("input is type chan int")
	default:
		fmt.Println("unrecognised type")
	}
}

func main() {
	v1 := 5
	v2 := "hello"
	v3 := true
	v4 := make(chan int)
	v5 := make(map[string]string)

	recogniseType(v1)
	recogniseType(v2)
	recogniseType(v3)
	recogniseType(v4)
	recogniseType(v5)
}
