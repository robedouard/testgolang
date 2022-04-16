package main

//testing
import "fmt"

var a, b, c int

func main() {
	x := []int{55, 20, 88, 100, 33, 95, 33, 84, 21, 93, 36, 9, 54, 3, 82, 4}
	b := x[0]
	for a = 0; a < len(x); a++ {
		fmt.Println(a, "we are evaluating: ", x[a], "against", b)
		if b <= x[a] {
			b = b
		} else {
			b = x[a]
		}
		c = b
	}
	fmt.Println("\nThe Lowest number in this Array is: ", c)
}
