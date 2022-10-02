package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to Maths in golang")

	var myInt int = 5

	fmt.Println(myInt)

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	n, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(n)
}
