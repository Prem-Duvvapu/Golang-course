package main

import "fmt"

const LogInToken string = "qwersdfgzxcv"; //public

func main() {
	var username string = "Prem";
	fmt.Println(username);
	fmt.Printf("Vraiable is of type: %T \n", username);

	var isLoggedIn bool=true;
	fmt.Println(isLoggedIn);
	fmt.Printf("Variable is of type: %T \n",isLoggedIn);

	var smallVal uint8=255;
	fmt.Println(smallVal);
	fmt.Printf("Variable is of type: %T \n",smallVal);

	var smallFloat float32=255.452234456789789567;
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n",smallFloat);

	//default values and aliases
	var anotherVariable int;
	fmt.Println(anotherVariable);
	fmt.Printf("Variable is of type: %T \n",anotherVariable);

	//implicit type
	var website="learncodeonline.in";
	fmt.Println(website);

	//no var style
	numberOfUser := 30000;
	fmt.Println(numberOfUser);

	fmt.Println(LogInToken);
	fmt.Printf("Variable is of type: %T \n",LogInToken);
}