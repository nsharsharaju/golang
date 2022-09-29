package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "My Text 1"

	file, err := os.Create("./harsha.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("The length of the file is", length)
	defer file.Close()
	readFile("./harsha.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("The data in the file", string(dataByte))
}
