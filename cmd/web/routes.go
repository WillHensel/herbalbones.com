package main

import "net/http"

func (app *application) getRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	mux.Handle("/", http.HandlerFunc(app.home))
	mux.Handle("/shop", http.HandlerFunc(app.shop))
	mux.Handle("/shop/buy-now", http.HandlerFunc(app.buyNowPost))
	mux.Handle("/contact", http.HandlerFunc(app.contact))

	return app.logRequest(mux)
}
