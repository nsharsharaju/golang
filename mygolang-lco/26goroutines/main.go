package main

import (
	"fmt"
	"net/http"
	"sync"
)

// func main() {
// 	fmt.Println("Welcome to Go Routines")
// 	go greeter("Hello")
// 	greeter("World")
// }

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(s)
// 	}
// }

var wg = sync.WaitGroup{}

var mut = sync.Mutex{}

var signals = []string{"test"}

func main() {
	fmt.Println("Welcome to Go Routines")
	websitesList := []string{
		"https://www.investwellonline.com",
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.fb.com",
		"https://www.lnmiit.ac.in",
	}

	for _, website := range websitesList {
		wg.Add(1)
		go getStatusCode(website)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS something is not good with endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
	wg.Done()
}
