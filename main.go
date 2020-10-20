package main

import (
	"fmt"
	"flag"
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
	fmt.Println("URL: ", url)
}