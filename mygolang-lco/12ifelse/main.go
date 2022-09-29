package main

import "fmt"

func main() {
	fmt.Println("If else in Go")

	loginCount := 23

	if loginCount < 10 {
		fmt.Println("Regular User")
	} else if loginCount > 10 {
		fmt.Println("Watch out")
	} else {
		fmt.Println("Exactly login count is 10")
	}

	if 9%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	if num := 3; num > 10 {
		fmt.Println("Number greater than 10")
	} else {
		fmt.Println("Number less than 10")
	}
}
