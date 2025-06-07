package main

import "fmt"

func main() {
	fmt.Println("Let's play with maps")

	languages := make(map[string]string)

	languages["JS"]="JavaScript"
	languages["RB"]="Ruby"
	languages["PY"]="Python"

	fmt.Println("List of all languages",languages)
	fmt.Println("RB shorts for",languages["RB"])

	delete(languages, "PY")
	fmt.Println("List of all languages",languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n",key,value)
	}
}