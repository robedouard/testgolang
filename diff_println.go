package main

import ("fmt")

func main(){
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(321325 * 424521)
	fmt.Println((true && false) || (false && true) || !(false && false))
}
