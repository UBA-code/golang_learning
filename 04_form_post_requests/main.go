package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	performFormPostRequest("https://jsonplaceholder.typicode.com/posts")
}

func performFormPostRequest(myUrl string) {
	formData := url.Values{}

	formData.Add("Name", "UBA")
	formData.Add("city", "meknes")
	formData.Add("job", "bluedove")

	response, err := http.PostForm(myUrl, formData)
	checkError(err)
	defer response.Body.Close()

	contentBytes, err := io.ReadAll(response.Body)
	checkError(err)
	fmt.Println(string(contentBytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
