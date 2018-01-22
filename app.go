package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	r, _ := http.NewRequest("GET", "http://google.com", nil)
	// r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	// r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(resp)
}
