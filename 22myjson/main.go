package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name		string		`json:"coursename"`
	Price		int			`json:"price"`
	Platform	string		`json:"website"`
	Password	string		`json:"-"`
	Tags		[]string	`json:"tags,omitempty"`
}


func main() {
	fmt.Println("Let's play with JSON")

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React JS", 2500, "namaste.dev", "react123", []string{"js", "react"}},
		{"Node JS", 3000, "namaste.dev", "react123", nil},
	}

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", finalJson)
}


func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "coursename": "React JS",
                "price": 2500,
                "website": "namaste.dev",
                "tags": [
                        "js",
                        "react"
                ]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}


	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key = %v and value = %v and value Type is: %T\n", k, v, v)
	}
}