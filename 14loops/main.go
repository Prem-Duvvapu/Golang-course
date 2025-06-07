package main

import "fmt"

func main() {
	fmt.Println("Let's play with loops")

	// days := []string{"Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"}

	// for d:=0;d<len(days);d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v \n",index,day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("value is %v \n",day)
	// }

	rogueValue := 1
	for rogueValue < 10 {
		// if rogueValue==3 {
		// 	rogueValue++
		// 	continue
		// }

		// if rogueValue==5 {
		// 	break
		// }

		if rogueValue==8 {
			goto lco
		}


		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

	lco:
		fmt.Println("Jumped to learncodeonline.in")
}