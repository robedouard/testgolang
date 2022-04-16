package main

import "fmt"

var loopcontrol, b, c int

func main() {
	x := []int{55, 20, 88, 100, 555, 33, 3, 95, 33, 84, 21, 93, 36, 9, 54, 82, 4}
	fmt.Println("This array has the following numbers: ", x[:])
	b := x[0]
	for loopcontrol = 0; loopcontrol < len(x); loopcontrol++ {
		fmt.Println(loopcontrol, "we are evaluating: ", x[loopcontrol], "against", b)
		if b >= x[loopcontrol] {
			b = b
		} else {
			b = x[loopcontrol]
		}
		c = b
	}
	fmt.Println("\nThe highest number in this Array is: ", c)
}
