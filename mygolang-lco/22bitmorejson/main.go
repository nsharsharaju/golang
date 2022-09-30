package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	fmt.Println("Welcome to Class on JSON")
	//encodedJson()
	decodeJson()
}

func encodedJson() {
	courses := []course{
		{"js", 10}, {"node", 12},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"name": "js",
		"price": 10
	}
	`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is VALID")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("The JSON is %#v\n", myCourse)
	} else {
		fmt.Println("Invalid JSON")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("The JSON is %#v\n", myOnlineData)

	for key, val := range myOnlineData {
		fmt.Println(key, val)
	}
}
