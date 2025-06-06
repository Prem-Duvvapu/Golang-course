package main

import "fmt"

func main() {
	fmt.Println("Let's play with pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is: ",ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of acutal pointer is: ",ptr)
	fmt.Println("Value in acutal pointer is: ",*ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}