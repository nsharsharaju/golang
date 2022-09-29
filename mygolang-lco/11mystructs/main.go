package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	harsha := User{"Harsha", "nsharsharaju@gmail.com", true, 26}

	fmt.Println(harsha)

	fmt.Printf("User details are %+v\n", harsha)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
