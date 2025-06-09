package main

import "fmt"
import "net/http"
import "io"

const url = "https://example.com"

func main() {
	fmt.Println("Let's play with web requests in golang")

	response, err := http.Get(url)

	if err!=nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n",response)

	defer response.Body.Close() //caller's responsiblity to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err!=nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("Content is: ",content)
}