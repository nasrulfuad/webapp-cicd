package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	port := ":2222"
	fmt.Printf("starting web server at http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
