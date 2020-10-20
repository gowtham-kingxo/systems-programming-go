package main

import (
	"fmt"
	"flag"
	"net/http"
	"io/ioutil"
)

func main() {
	var urlFlag = flag.String("url", "NO URL PROVIDED", "Make HTTP request to the given URL and display the response")
	var profileFlag = flag.Int("profile", -1, "Make HTTP requests to the profile url for that many times as the given number")
	flag.Parse()

	if (*urlFlag != "NO URL PROVIDED") {
		makeHTTPRequest(*urlFlag)

	} else if (*profileFlag != -1) {
		fmt.Println("Profile flag used: ", *profileFlag)

	} else {
		fmt.Println("No flags used...")
	}
}

func makeHTTPRequest(url string) {
	resp, err := http.Get("http://127.0.0.1:8787/links")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(string(body))
}