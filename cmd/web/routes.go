package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) getRoutes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	dynamic := alice.New(app.sessionManager.LoadAndSave, app.noSurf)

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/shop", app.shop)
	router.HandlerFunc(http.MethodGet, "/shop/buy-now/:id", app.buyNow)
	router.Handler(http.MethodGet, "/contact", dynamic.ThenFunc(app.contact))
	router.Handler(http.MethodPost, "/contact", dynamic.ThenFunc(app.contactPost))

	standard := alice.New(app.logRequest, app.secureHeaders)

	return standard.Then(router)
}
