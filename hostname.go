package main

import "fmt"
import "os"
import "github.com/fatih/color"

func main() {
	red := color.New(color.FgRed).PrintlnFunc()
	me, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	} else {
		red("hello\n")
		red(me, "\n")
	}

}
