package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 30
	var b string = "Hello World"
	var c float32 = 33.888
	var d bool = true
	fmt.Println(a, "   variable a is a:   ", reflect.TypeOf(a))
	fmt.Println(b, "   variable b is a:   ", reflect.TypeOf(b))
	fmt.Println(c, "   variable c is a:   ", reflect.TypeOf(c))
	fmt.Println(d, "   variable d is a:   ", reflect.TypeOf(d))
}
