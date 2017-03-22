package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/page/", PageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Hello world :) <a href=\"page/\">continue</a>")
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "12345...")
}
