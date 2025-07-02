package main

import (
	// "context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Let's play with web verb requests")

	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}


func PerformGetRequest() {
	const myurl="http://example.com"

	response, err := http.Get(myurl)
	if err!=nil {
		panic(err)
	}

	defer response.Body.Close();

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ",byteCount)
	fmt.Println(responseString.String())
}

func PerformPostRequest() {
	const myurl="https://httpbin.org/post"

	//json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 500,
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl="https://postman-echo.com/post"

	//form data
	data := url.Values{}
	data.Add("firstname", "prem")
	data.Add("firstname", "sai")
	data.Add("lastname","duvvapu")

	response, err := http.PostForm(myurl, data)
	if err!=nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}