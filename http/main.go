package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// To get the actual body read this document: https://golang.org/pkg/net/http/#Response
	fmt.Println(resp)
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}

// link for standard documenation  https://golang.org/pkg/net/http/, https://golang.org/pkg/net/http/#Get
