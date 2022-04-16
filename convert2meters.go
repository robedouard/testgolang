package main

import "fmt"

func main() {
	var answer float64

	fmt.Println("Please enter how many feet you would like to convert to meters:")

	fmt.Scanf("%f", &answer)

	var output float64

	output = answer * 0.3048

	fmt.Println(output)
}
