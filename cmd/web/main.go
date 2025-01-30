package main

import (
	"net/http"

	"github.com/a-h/templ"
	"web.herbalbones.com/ui/pages"
)

func main() {
	http.Handle("/", templ.Handler(pages.Home()))
	http.ListenAndServe("localhost:4000", nil)
}
