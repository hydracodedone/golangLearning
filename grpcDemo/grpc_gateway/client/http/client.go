package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	type request struct {
		Id int32 `json:"id,omitempty"`
	}
	data := request{Id: 1}
	dataBytes, err := json.Marshal(data)
	fmt.Println(dataBytes)
	if err != nil {
		panic(err)
	}
	bodyReader := bytes.NewReader(dataBytes)
	resp, err := http.Post("http://localhost:9001/http_request", "application/json", bodyReader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status code:[%v] body:[%v]", resp.StatusCode, string(body))
}
