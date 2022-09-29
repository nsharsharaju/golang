package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to class on Switch case")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1

	fmt.Println("The dice number is", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6")
	default:
		fmt.Println("What was that!")

	}
}
