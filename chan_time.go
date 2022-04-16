package main

import (
	"fmt"
	"time"
)

func receiver(cvariable chan string) {
	for msg := range cvariable {
		fmt.Println(msg)
	}
}

//hello this is a test
func main() {
	messages := make(chan string, 3)
	messages <- "hello"
	messages <- "world"
	messages <- "testing"
	close(messages)
	fmt.Println("Pushed messages onto channel with no receivers")
	time.Sleep(time.Second * 1)
	receiver(messages)
}
