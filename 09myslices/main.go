package main

import "fmt"
import "sort"

func main() {
	fmt.Println("Let's play with slices")

	var fruitList = []string{"Apple","Banana","Mango"}
	fmt.Printf("Type of fruitList is: %T\n",fruitList)

	fruitList1 := append(fruitList,"Pineapple")
	fmt.Println(fruitList1)

	fruitList2 := append(fruitList[1:3])
	fmt.Println(fruitList2)


	highScores := make([]int, 4)
	highScores[0]=456
	highScores[1]=973
	highScores[2]=237
	highScores[3]=612

	fmt.Println("HighScores: ",highScores)

	// highScores[4]=384 //gives panic: runtime error: index out of range [4] with length 4

	highScores = append(highScores, 384, 529) //this will work
	fmt.Println("HighScores: ",highScores)

	sort.Ints(highScores)
	fmt.Println("sorted HighScores: ",highScores)
	fmt.Println(sort.IntsAreSorted(highScores))


	//how to remove a value from slice based on index
	var courses=[]string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(courses)

	var index int=2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}