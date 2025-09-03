package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var justString string

func createHugeString(size int) string {
	var str strings.Builder
	for range size {
		// строка состоит из цифр 0-9
		str.WriteString(fmt.Sprintf("%d", rand.Intn(10)))
	}
	return str.String()
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
