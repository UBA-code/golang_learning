package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/4"

func main() {
	response, err := http.Get(url) //? we did a get request to the constant url
	checkError(err)                //? check for errors (network error, ...)
	if response.StatusCode != 200 {
		panic("response code : " + response.Status + "!") //* if the response status is not 200 OK we exit
	}
	defer response.Body.Close() //! it's important to close the response body, it can cause consume more memory and resources
	dataBytes, err := io.ReadAll(response.Body)
	checkError(err)
	fmt.Println(string(dataBytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
