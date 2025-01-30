package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	mux.Handle("/", http.HandlerFunc(home))
	mux.Handle("/shop", http.HandlerFunc(shop))
	mux.Handle("/contact", http.HandlerFunc(contact))

	http.ListenAndServe("localhost:4000", mux)
}
