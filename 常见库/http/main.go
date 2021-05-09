package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func demo() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Printf("request failed: %v\n", err)
	}
	defer func() {
		resp.Body.Close()
	}()
	body, err := os.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("responses failed: %v\n", err)
	}
	fmt.Printf("the response body is: %v\n", string(body))
}
func main() {
	demo()
}
