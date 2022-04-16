package main

import "fmt"

var value, ismult, summult int

func main() {
	for value = 0; value < 1000; value++ {
		if value%3 != 0 && value%5 != 0 {
		} else {
			ismult = value
			summult += ismult
		}
	}
	fmt.Println(summult)
}
