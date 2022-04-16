package main

import (
	"fmt"
)

func main() {
	var userInput float64 
	fmt.Println("Please enter the temperature you would like to convert to Celsius")
	fmt.Scanf("%f", &userInput)
	var output float64 
	output = (( userInput - 32 ) * 5/9 )
	fmt.Println( output )
}
