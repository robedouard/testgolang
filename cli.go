package main

import (
	"flag"
	"fmt"
)

func main() {
	var cmd string
	flag.StringVar(&cmd, "o", cmd, `option can be either "add" or "remove"`)

	flag.Parse()
	switch cmd {
	case "add":
		fmt.Println("Hello")
	case "remove":
		fmt.Println("bye")
	default:
		flag.Usage()
	}
}
