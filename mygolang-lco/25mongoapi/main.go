package main

import (
	"fmt"
	"net/http"

	"github.com/nsharsharaju/mongoapi/router"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server Getting started...")
	http.ListenAndServe(":4000", r)
}
