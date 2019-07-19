package main

import (
	"fmt"
	"net/http"
)

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
	fmt.Fprintf(w, "<p>Go is fast</p>")
	fmt.Fprintf(w, "<p>%s %s people</p>", "Very", "<strong>happy</strong>")
}

func main() {
	http.HandleFunc("/", start)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":57855", nil)
}
