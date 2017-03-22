package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/oven/", OvenHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<html>
<head>
<title>Welcome</title>
</head>
<body>
<form action="/oven/" method="post">
<label>What kind of pizza do you like?</label>
<input type="text" name="my_pizza">
<input type="submit">
</form>
</body>
</html>`)
}

func OvenHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pizza := r.Form.Get("my_pizza")

	if pizza == "" {
		http.Error(w, "Error: No pizza specified!", http.StatusInternalServerError)
		return
	}

	s := fmt.Sprintf("Serving <strong>%s pizza</strong>, just for you!", html.EscapeString(pizza))

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, s)
}
