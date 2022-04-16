package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Enter a number to multiply by psuedo random number:")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * rand.Float64()
	fmt.Println(output)
}
