package main

import (
	"fmt"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", "https://en.wikipedia.org/wiki/Special:Random", nil)

	resp, err := client.Do(req)

	// 	resp, err := http.Get("https://en.wikipedia.org/wiki/Special:Random")
	check(err)
	// 	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header["Location"][0])
}
