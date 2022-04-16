package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  uint
}

var molten string

// sender sends message and time to a channel
func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I have a message"
		<-t.C
	}

}

func main() {
	messages := make(chan string)

	stop := make(chan bool)

	go sender(messages)

	go func() {
		time.Sleep(time.Second * 4)
		fmt.Println("Times up")
		stop <- true
		//TODO
	}()

	for {
		select {
		case <-stop:
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
