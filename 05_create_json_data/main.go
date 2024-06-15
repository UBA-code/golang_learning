package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"` //? change the name of the attribute in case of json
	Price    float64
	Password string   `json:"-"`              //? didn't show password
	Tags     []string `json:"tags,omitempty"` //? didn't showed in case of null
}

func main() {
	createJson()
}

func createJson() {
	jsonData := []course{
		{Name: "ReactJs", Price: 19.99, Password: "test123", Tags: []string{"Js", "Ts", "React"}},
		{Name: "C++", Price: 199.99, Password: "c++master", Tags: []string{"c", "c++"}},
		{Name: "Docker", Price: 999, Password: "docker", Tags: nil},
	}

	jsonContentBytes, err := json.Marshal(jsonData)
	// jsonContentBytes, err := json.MarshalIndent(jsonData, "", "\t") //* format it to be more readable
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonContentBytes))
}
