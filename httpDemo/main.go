package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func server() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8091", nil)
	fmt.Println(12)
}

func client() {
	data := []byte(`{"name":"hydra"}`)

	resp, err := http.Post("http://localhost:8091", "application/json", bytes.NewBuffer(data))
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		respData, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(respData))
	}
}
func main() {

}
