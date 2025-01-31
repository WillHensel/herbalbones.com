package main

import (
	"net/http"

	"web.herbalbones.com/ui/pages"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	pages.Home().Render(r.Context(), w)
}

func (app *application) shop(w http.ResponseWriter, r *http.Request) {
	pages.Shop().Render(r.Context(), w)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {

	pages.Contact().Render(r.Context(), w)
}
