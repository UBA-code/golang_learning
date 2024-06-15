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
	getJson()
}

func getJson() {
	jsonData := []byte(`
		{
			"courseName": "ReactJs",
			"Price": 19.99,
			"tags": [
							"Js",
							"Ts",
							"React"
			]
		}
	`)

	if json.Valid(jsonData) { //? check if json data is a valid json

		var dataObj course

		err := json.Unmarshal(jsonData, &dataObj) //? encoding the json data to go object
		if err != nil {
			panic(err)
		}

		fmt.Printf("%#v\n", dataObj)

		//* we can also get the json as key value map
		jsonDataAsMap := map[string]any{}

		err = json.Unmarshal(jsonData, &jsonDataAsMap)
		//? we can check if the Unmarsharll failed

		fmt.Printf("%#v\n", jsonDataAsMap)

	} else {
		fmt.Println("JSON DATA IS NOT VALID")
	}
}
