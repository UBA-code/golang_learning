package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("tempFile") //? for creating new file
	checkError(err)
	defer file.Close() //* this line will run after the function finished

	len, err := io.WriteString(file, "This is a test content") //? for writing strings to files
	checkError(err)
	fmt.Println("content length: ", len)
	readFile(file.Name())
}

//? this function read data from the filepath argument
func readFile(filePath string)  {
	dataBytes, err := os.ReadFile(filePath) //? for reading data for a file, return a bytes slice
	checkError(err)
	fmt.Println(string(dataBytes))
}

//? function for checking error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}