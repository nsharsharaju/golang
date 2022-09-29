package main

import "fmt"

func main() {
	fmt.Println("Welcome to Video for Loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for index := range days {
		fmt.Println(days[index])
	}

	for _, day := range days {
		fmt.Println(day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		fmt.Println(rougueValue)
		rougueValue++
	}
}
