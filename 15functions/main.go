package main

import "fmt"

func main() {
	fmt.Println("Let's play with functions in golang")
	greeter()

	result := adder(3,5)
	fmt.Println("Result is: ",result)

	proResult, proMessage := proAdder(1,4,6,7,9,8)
	fmt.Println("pro Result is: ",proResult)
	fmt.Println("pro Message is: ",proMessage)
}

func greeter() {
	fmt.Println("Namaste golang")
}

func adder(a int,b int) int {
	return a+b
}

func proAdder(values ...int) (int,string) {
	total := 0

	for _,value := range values {
		total += value
	}

	return total, "Hi proResult"
}