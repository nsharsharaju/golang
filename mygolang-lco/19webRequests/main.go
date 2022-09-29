package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to class on Web Requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The type of the response is %T", response)
	defer response.Body.Close() // Responsibility of the developer to close the connection

	dataByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataByte)

	fmt.Println(content)

}
