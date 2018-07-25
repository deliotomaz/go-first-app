package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewAppRouter()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	log.Fatal(http.ListenAndServe(":3000", router))

}
