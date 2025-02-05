package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"web.herbalbones.com/ui/components"
	"web.herbalbones.com/ui/pages"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	pages.Home().Render(r.Context(), w)
}

func (app *application) shop(w http.ResponseWriter, r *http.Request) {

	catalogItems, err := app.services.squareService.CatalogList()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	model := []components.CatalogItemCardViewModel{}

	for _, i := range catalogItems {
		image := ""
		if len(i.ImageUrls) > 0 {
			image = i.ImageUrls[0]
		}
		price := 0
		id := ""
		// TODO support multiple variations
		if len(i.Variations) > 0 {
			price = i.Variations[0].Price
			id = i.Variations[0].Id
		}
		item := components.CatalogItemCardViewModel{
			Id:              id,
			PrimaryImageUrl: image,
			Price:           price,
			Name:            i.Name,
		}
		model = append(model, item)
	}

	pages.Shop(model).Render(r.Context(), w)
}

func (app *application) buyNow(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	if len(strings.TrimSpace(id)) == 0 {
		app.notFound(w)
		return
	}

	uri, err := app.services.squareService.GetSingleItemPaymentLink(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Info("Buy Now", "item_id", id, "uri", uri)

	http.Redirect(w, r, uri, http.StatusSeeOther)
}

func (app *application) buyNowPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	itemId := r.PostForm.Get("item_id")
	if len(strings.TrimSpace(itemId)) == 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	uri, err := app.services.squareService.GetSingleItemPaymentLink(itemId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Info("Buy Now", "item_id", itemId, "uri", uri)

	http.Redirect(w, r, uri, http.StatusSeeOther)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {

	pages.Contact().Render(r.Context(), w)
}
