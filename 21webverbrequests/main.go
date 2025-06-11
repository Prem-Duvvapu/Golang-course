package main

import "fmt"
import "net/http"
import "io"
import "strings"

func main() {
	fmt.Println("Let's play with web verb requests")

	PerformGetRequest()
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