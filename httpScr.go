package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	temp_server, _ := http.Get("http://blank.com")
	input, _ := ioutil.ReadAll(temp_server.Body)
	string_it := string(input)
	fmt.Println(string_it)
	temp_server.Body.Close()

}
