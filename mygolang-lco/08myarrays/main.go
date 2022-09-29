package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array Class")

	var fruitsList [5]string

	fruitsList[0] = "Apple"
	fruitsList[1] = "Banana"
	fruitsList[3] = "Peach"

	fmt.Println("The fruit list is", fruitsList)
}
