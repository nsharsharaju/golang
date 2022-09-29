package main

import "fmt"

func main() {
	fmt.Println("We are now exploring on functions")
	greeter()

	result := adder(1, 2)
	fmt.Println("The result of the function is", result)

	result = proAdder(1, 2, 3)
	fmt.Println("The result of the function is", result)

}

func greeter() {
	fmt.Println("Namastey from golang")
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}
