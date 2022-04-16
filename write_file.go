package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bleh := []byte("wan, wan, wan")
	err := ioutil.WriteFile("/tmp/myfile.txt", bleh, 0644)
	if err != nil {
		fmt.Println(err)
	}
	grah, _ := os.Create("/tmp/break")
	fmt.Println(grah)
	f, err := ioutil.ReadFile("/tmp/myfile.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("When babies cry, their bytes sound like", f)
	fmt.Println("When babies cry it sounds like", string(f))
	os.Remove("/tmp/myfile.txt")
	os.Remove("/tmp/break")
}
