package main

import (
	"fmt"
	"math/big"
)

func sum(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {
	a := new(big.Int).Exp(big.NewInt(3), big.NewInt(50), nil) // 2^70
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(70), nil)
	fmt.Println("sum: ", sum(a, b))
	fmt.Println("sub: ", sub(a, b))
	fmt.Println("mul: ", mul(a, b))
	fmt.Println("div: ", div(a, b))

}
