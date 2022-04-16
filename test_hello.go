package main

import (
	"fmt"
)

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(3 * 4)
	fmt.Println((true && false) || (false && true) || !(false && false))

	var x string = "Hello"
	fmt.Println(x)
	x = x + "world"
	fmt.Println(x)

}
