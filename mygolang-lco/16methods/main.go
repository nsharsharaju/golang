package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	harsha := User{"Harsha", "nsharsharaju@gmail.com", true, 26}

	fmt.Println(harsha)

	fmt.Printf("User details are %+v\n", harsha)

	harsha.GetStatus()
	harsha.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("The status of the users is", u.Status)
}

func (u User) NewMail() {
	u.Email = "test2@go.dev"
	fmt.Println("The new email of the users is", u.Email)
}
