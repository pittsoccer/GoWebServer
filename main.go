package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(response http.ResponseWriter, request *http.Request) {
	// output user
	fmt.Fprintf(response, "<h1>Hello Bryce!</h1>")

	// output time
	fmt.Fprintf(response, "<h3> The time is: ")
	fmt.Fprintf(response, time.Now().Format(time.RFC1123))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}
