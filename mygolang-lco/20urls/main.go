package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn/?courseName=reactJS&paymentid=kdskadhksajh"

func main() {
	fmt.Println("Welcome to class on URLs")

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Println(qparams)

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/learn",
		RawPath: "user=harsha",
	}

	newUrl := partsOfUrl.String()

	fmt.Println(newUrl)
}
