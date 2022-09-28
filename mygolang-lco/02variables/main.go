package main

import "fmt"

const LoginToken = "hdkjshdkshdk"

func main() {
	var username string = "nsharsharaju"
	fmt.Println(username)
	fmt.Printf("The type of the variable is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of the variable is %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("The type of the variable is %T \n", smallInt)

	var smallFloat float64 = 255.80283
	fmt.Println(smallFloat)
	fmt.Printf("The type of the variable is %T \n", smallFloat)

	//default values and some aliasas
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("The type of the variable is %T \n", anotherVariable)

	//implicit types
	var website = "nsharsharaju.com"
	fmt.Println(website)

	//no var style
	numberOfUsers := 500
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("The type of the variable is %T\n", LoginToken)
}
