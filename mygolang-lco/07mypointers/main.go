package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("The value in the pointer is", ptr)

	var number int = 23

	var ptr = &number

	fmt.Println("The value in the pointer is", ptr)

	*ptr = *ptr + 2

	fmt.Println("The value in the number is", *ptr)
}
