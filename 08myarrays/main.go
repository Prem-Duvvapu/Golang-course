package main

import "fmt"

func main() {
	fmt.Println("Let's play with arrays")

	var fruits [5]string
	fruits[0]="apple"
	fruits[1]="banana"
	fruits[2]="mango"
	fruits[4]="pineapple"

	fmt.Println("Fruits list",fruits)
	fmt.Println("Fruits list length",len(fruits))

	var vegetables=[4]string{"potato", "carrot","brinjal"}
	fmt.Println("vegetables list length",len(vegetables))
}