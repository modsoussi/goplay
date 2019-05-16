package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/protected", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "protected")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `Hello, World!`)
	})

	http.ListenAndServe(":8080", nil)
}
