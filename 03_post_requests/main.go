package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Performing Post request with golang")

	performPostRequest("https://jsonplaceholder.typicode.com/posts")
}

func performPostRequest(url string) {
	json := strings.NewReader(`
	{
		"name":"UBA",
		"city":"Meknes",
		"job":"bluedove"
	}
	`)
	reponse, err := http.Post(url, "application/json", json)
	checkError(err)
	defer reponse.Body.Close()

	contentbytes, err := io.ReadAll(reponse.Body)
	checkError(err)

	fmt.Println(string(contentbytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
