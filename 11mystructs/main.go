package main

import "fmt"

func main() {
	fmt.Println("Let's play with structs")
	//no inheritance in golang; No super or parent

	prem := User{"Prem","prem123@learning.in",true,23}
	fmt.Println(prem)
	fmt.Printf("prem details are: %+v\n", prem)
	fmt.Printf("Name is: %v and Email is: %v\n", prem.Name, prem.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}