package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getRoutes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/shop", app.shop)
	router.HandlerFunc(http.MethodGet, "/shop/buy-now/:id", app.buyNow)
	// mux.Handle("/shop/buy-now", http.HandlerFunc(app.buyNowPost))
	router.HandlerFunc(http.MethodGet, "/contact", app.contact)

	return app.logRequest(router)
}
