package main

import "fmt"
import "bufio"
import "os"

func main() {
	welcome := "Welcome to user input";
	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter the rating for our Pizza:");

	//comma ok  or err ok
	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for rating, ", input);
	fmt.Printf("Type of this rating is %T ", input);
}