package main

import "fmt"
import "net/url"

const myurl = "https://httpbin.org/get?coursename=reactjs&paymentid=asdfqwer"

func main() {
	fmt.Println("Let's play with URLs in golang")

	fmt.Println(myurl)

	//parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	fmt.Println()

	qparams := result.Query()
	fmt.Printf("Type of qparams is %T\n",qparams)
	fmt.Println(qparams)

	for key, value := range qparams {
		fmt.Println(key,value)
	}

	partsOfUrl := &url.URL {
		Scheme: "https",
		Host: "httpbin.org",
		Path: "/get",
		RawQuery: "coursename=golang",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}