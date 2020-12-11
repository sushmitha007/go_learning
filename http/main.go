package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// To get the actual body read this document: https://golang.org/pkg/net/http/#Response
	// fmt.Println(resp)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Using Writer interface
	// io.Copy(os.Stdout, resp.Body)

	//Using custom Write
	lw:=logWriter{}
	io.Copy(lw,resp.Body)
}

// Custom Write Implementation
func(logWriter) Write(bs []byte)(int,error){
	fmt.Println(string(bs))
	fmt.Println("just wrote these many bytes:",len(bs))
	return len(bs), nil

}
// link for standard documenation  https://golang.org/pkg/net/http/, https://golang.org/pkg/net/http/#Get
