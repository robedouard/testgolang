package main

import "fmt"

// This import is for the hidden user input
import "golang.org/x/crypto/ssh/terminal"

func main() {

	var userInput string
	fmt.Println("Hello please enter a value :\n")
	fmt.Scanf("%s", &userInput)
	fmt.Println("Here is your input from scanf :", userInput)

	// Hide user input
	fmt.Println("Enter hidden input :\n")
	hiddenPass, _ := terminal.ReadPassword(0)
	fmt.Printf("Here is your input from terminal.ReadPassword : %s", hiddenPass)

}
