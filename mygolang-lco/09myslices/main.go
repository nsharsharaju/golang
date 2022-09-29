package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("The fruit list is", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("The fruit list is", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("The fruit list is", fruitList)

	//using make

	highScores := make([]int, 4)

	highScores[0] = 4
	highScores[1] = 2
	highScores[2] = 1
	highScores[3] = 3

	highScores = append(highScores, 6, 5)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"js", "html", "css", "ruby", "python"}

	index := 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
