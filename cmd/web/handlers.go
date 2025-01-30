package main

import (
	"net/http"

	"web.herbalbones.com/ui/pages"
)

func home(w http.ResponseWriter, r *http.Request) {

	pages.Home().Render(r.Context(), w)
}

func shop(w http.ResponseWriter, r *http.Request) {
	pages.Shop().Render(r.Context(), w)
}

func contact(w http.ResponseWriter, r *http.Request) {

	pages.Contact().Render(r.Context(), w)
}
