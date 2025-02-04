// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

type CatalogItemCardViewModel struct {
	Id              string
	PrimaryImageUrl string
	Price           int
	Name            string
}

func CatalogItemCardComponent(model CatalogItemCardViewModel) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<style>\r\n\t\t.product-card {\r\n\t\t\tdisplay: flex;\r\n\t\t\tflex-direction: column;\r\n\t\t\talign-items: center;\r\n\t\t}\r\n\r\n\t\t.product-card * {\r\n\t\t\tcolor: white;\r\n\t\t}\r\n\r\n\t\t.product-card-name {\r\n\t\t\tfont-size: 2em;\r\n\t\t\tmargin-block: 0.25rem;\r\n\t\t\ttext-align: center;\r\n\t\t}\r\n\r\n\t\t.product-card-price {\r\n\t\t\ttext-align: center;\r\n\t\t\tmargin-block: 0\r\n\t\t}\r\n\r\n\t\t.product-card-image {\r\n\t\t\twidth: 100%;\r\n\t\t\theight: 25rem;\r\n\t\t\taspect-ratio: square;\r\n\t\t\tobject-fit: cover;\r\n\t\t\tobject-position: center;\r\n\t\t}\r\n\r\n\t\t.buy-button {\r\n\t\t\ttext-align: center;\r\n\t\t\ttext-decoration: none;\r\n\t\t}\r\n\r\n\t\t.buy-button * {\r\n\t\t\tcolor: #6b6a65;\r\n\t\t}\r\n\r\n\t\t.buy-button > form{\r\n\t\t\twidth: 10rem;\r\n\t\t\theight: 2.5rem;\r\n\t\t\tbackground-color: white;\r\n\t\t\tborder-radius: 3px;\r\n\t\t\ttext-align: center;\r\n\t\t\tmargin-top: 0.5rem;\r\n\t\t\tdisplay: flex;\r\n\t\t\talign-items: center;\r\n\t\t\tjustify-content: center;\r\n\t\t}\r\n\r\n\t\t.buy-button > form > button {\r\n\r\n\t\t}\r\n\t</style><div class=\"product-card\"><img class=\"product-card-image\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(model.PrimaryImageUrl)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/components/catalog_item_card.templ`, Line: 69, Col: 61}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" alt=\"\"><h3 class=\"product-card-name\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(model.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/components/catalog_item_card.templ`, Line: 70, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</h3><p class=\"product-card-price\">$")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", model.Price/100))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/components/catalog_item_card.templ`, Line: 71, Col: 71}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, ".00</p><a class=\"buy-button\" href=\"#\"><form action=\"/shop/buy-now\"><input type=\"hidden\" name=\"item_id\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(model.Id)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/components/catalog_item_card.templ`, Line: 74, Col: 56}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\"> <button type=\"submit\">Buy Now</button></form></a></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
