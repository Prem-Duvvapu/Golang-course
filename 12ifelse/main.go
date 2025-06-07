package main

import "fmt"

func main() {
	fmt.Println("Let's play with if else")

	loginCount := 3
	var result string

	if loginCount<10 {
		result = "Regular user"
	} else if loginCount>10 {
		result = "Premium user"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)
	

	if num:=3; num%2==1 {
		fmt.Println("Nums is odd")
	} else {
		fmt.Println("Num is even")
	}

	// if err!=nil {

	// }
}