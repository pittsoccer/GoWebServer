package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(response, "Hello %s!", request.URL.Path[1:])
	fmt.Fprintf(response, "Hello Bryce!")
	fmt.Fprintf(response, "\n")
	fmt.Fprintf(response, time.Now().Format(time.RFC1123))
}

func main() {
	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s!\n", s)

	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}
