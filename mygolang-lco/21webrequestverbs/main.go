package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Requests")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("The status code is", response.StatusCode)
	fmt.Println("The Content length is", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	var responseString strings.Builder
	// fmt.Println(string(content))

	byteCount, err := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}
