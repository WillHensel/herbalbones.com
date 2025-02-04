package main

import (
	"net/http"

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
		if len(i.Variations) > 0 {
			price = i.Variations[0].Price
		}
		item := components.CatalogItemCardViewModel{
			Id:              i.Id,
			PrimaryImageUrl: image,
			Price:           price,
			Name:            i.Name,
		}
		model = append(model, item)
	}

	pages.Shop(model).Render(r.Context(), w)
}

func (app *application) buyNowPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	itemId := r.PostForm.Get("item_id")
	uri, err := app.services.squareService.GetSingleItemPaymentLink(itemId)
	if err != nil {
		return
	}

	http.Redirect(w, r, uri, http.StatusSeeOther)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {

	pages.Contact().Render(r.Context(), w)
}
