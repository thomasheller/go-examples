package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/numbers/", NumbersHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `Hello world :) <a href="numbers/">continue</a>`)
}

func NumbersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "67890...")
}
