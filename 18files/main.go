package main

import "fmt"
import "os"
import "io"

func main() {
	fmt.Println("Let's play with defer in golang")
	content := "This needs to go in file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file,content)
	checkNilErr(err)
	fmt.Println("Length is",length)
	defer file.Close()

	readLine("./mylcogofile.txt")
}

func readLine(filename string) {
	databyte,err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err!=nil {
		panic(err)
	}
}