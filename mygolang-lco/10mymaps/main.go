package main

import "fmt"

func main() {
	fmt.Println("Welcome to video on maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)

	fmt.Println(languages["JS"])

	delete(languages, "JS")

	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println(key, value)
	}
}
