package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating", input)
	fmt.Printf("The type of input is %T\n", input)

}
