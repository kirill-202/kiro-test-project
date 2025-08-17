package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	
	fmt.Fprintf(w, "<html><body><h1>Hello World App</h1><pre>")
	for i := 1; i <= 25; i++ {
		fmt.Fprintf(w, "Hello World\n")
	}
	fmt.Fprintf(w, "</pre></body></html>")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}