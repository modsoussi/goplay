package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	log.Println("Server starting")

	http.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request){
		
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `Hello, World!`)
	})

	http.ListenAndServe(":8080", nil)
}
