package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Requests")
	PerformGetRequest()
	PerformPostRequest()
	PerformPostFormRequest()
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

func PerformPostRequest() {
	const myurl string = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
	 	{
		"name": "harsha"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl string = "http://localhost:8000/postform"

	//form-data

	form := url.Values{}

	form.Add("name", "harsha")

	response, err := http.PostForm(myurl, form)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
