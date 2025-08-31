package main

import (
	"fmt"
	"log"
)

// examples:
// number=10, i = 3 заменить на 1 -> 14 (1010 -> 1100 = 14)
// number=10, i = 2 заменить на 0 -> 8	(1010 -> 1000 = 8)

func changeToZero(number int64, i int) int64 {
	var n int64 = 1 << (i - 1)
	return number &^ n
}
func changeToOne(number int64, i int) int64 {
	var n int64 = 1 << (i - 1)
	fmt.Println(n)
	return number | n

}

func main() {
	var number int64
	var i, val int
	fmt.Println("input number")
	fmt.Scan(&number)
	if number < 0 {
		log.Fatalln("incorrect number")
	}
	fmt.Println("input i")
	fmt.Scan(&i)
	if i <= 0 || i > 64 {
		log.Fatalln("incorrect i")
	}
	fmt.Println("input 0 or 1")
	fmt.Scan(&val)
	if val != 0 && val != 1 {
		log.Fatalln("incorrect input")
	}
	var newNumber int64
	if val == 0 {
		newNumber = changeToZero(number, i)
	} else {
		newNumber = changeToOne(number, i)
	}
	fmt.Printf("new number: %d\n", newNumber)

}
