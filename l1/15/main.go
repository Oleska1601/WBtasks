package main

import (
	"fmt"
	"math/rand"
)

var justString string

func createHugeString(size int) string {
	str := ""
	for range size {
		// строка состоит из цифр 0-9
		str += fmt.Sprintf("%d", rand.Intn(10))
	}
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	// создаем новый массив байтов и копируем в него нужное содержимое
	// преобразуем к строке
	justString = string([]byte(v[:100]))
	fmt.Println(justString)
	fmt.Println(len(justString), cap([]byte(justString)))
}

func main() {
	someFunc()
}
