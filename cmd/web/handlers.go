package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"web.herbalbones.com/ui/components"
	"web.herbalbones.com/ui/pages"
)

type templateData struct {
	Form      any
	Flash     string
	CSRFToken string
}

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

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	vm := pages.ContactViewModel{
		CSRFToken: data.CSRFToken,
		Flash:     data.Flash,
	}

	pages.Contact(vm).Render(r.Context(), w)
}

func (app *application) contactPost(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	vm := pages.ContactViewModel{
		CSRFToken: data.CSRFToken,
		Flash:     data.Flash,
		Errors:    make(map[string]string),
	}

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	message := r.PostForm.Get("message")

	vm.Name = name
	vm.Email = email
	vm.Message = message

	if len(strings.TrimSpace(name)) == 0 {
		vm.Errors["name"] = "Name field is required"
	}

	if len(strings.TrimSpace(email)) == 0 {
		vm.Errors["email"] = "Email field is required"
	}

	if len(strings.TrimSpace(message)) == 0 {
		vm.Errors["message"] = "Message field is required"
	}

	if len(vm.Errors) > 0 {
		pages.Contact(vm).Render(r.Context(), w)
		return
	}

	err = app.services.mailjetService.SendContactMessage(name, email, message)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Successfully sent message. We'll get back to you as soon as we can.")

	app.logger.Info("New contact request received", "name", name, "email", email, "message", message)

	http.Redirect(w, r, "/contact", http.StatusSeeOther)
}
