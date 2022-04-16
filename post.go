package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	userpass, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Enter pass :\n")

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{Jar: jar}
	response, err := client.PostForm("https://postsite.com/wp-login.php", url.Values{"log": {"username_here"}, "pwd": {string(userpass)}})
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", string(body))
}
