package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://jsonplaceholder.typicode.com/todos/4?name=UBA&city=Meknes&job=bluedove"

func main()  {
	result, err := url.Parse(myUrl) //? for parsing the url
	checkError(err)
	//? print all url information
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)


	queryMap := result.Query() //? return a map with key value of the url queries
	for name, val := range queryMap {
		fmt.Println(name, ":", val)
	}

	createdUrl := &url.URL{
		Scheme: "https",
		Host: "jsonplaceholder.typicode.com",
		Path: "/todos/2",
		RawQuery: "name=UBA&city=Meknes&job=bluedove",
	}

	fmt.Println(createdUrl.String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}