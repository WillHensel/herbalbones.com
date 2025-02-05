// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"web.herbalbones.com/ui"
)

func Layout(content templ.Component) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width\"><title>Herbal Bones</title><link rel=\"preconnect\" href=\"https://fonts.googleapis.com\"><link rel=\"preconnect\" href=\"https://fonts.gstatic.com\" crossorigin><link href=\"https://fonts.googleapis.com/css2?family=Birthstone&amp;family=Noto+Serif:ital,wght@0,100..900;1,100..900&amp;display=swap\" rel=\"stylesheet\"><link rel=\"stylesheet\" href=\"/static/site.css\"></head><body><style>\r\n\t\t\t\tbody {\r\n\t\t\t\t\twidth: 100%;\r\n\t\t\t\t\theight: 100%;\r\n\t\t\t\t}\r\n\r\n\t\t\t\theader {\r\n\t\t\t\t\tposition: relative;\r\n\t\t\t\t}\r\n\r\n\r\n\r\n                footer {\r\n                    height: 3rem;\r\n                    background-color: #38383b;\r\n\t\t\t\t\tdisplay: flex;\r\n\t\t\t\t\tjustify-content: space-between;\r\n\t\t\t\t\talign-items: center;\r\n\t\t\t\t\tpadding-inline: 1.5rem;\r\n                }\r\n\r\n\t\t\t\tfooter p {\r\n\t\t\t\t\tfont-size: 14px;\r\n\t\t\t\t}\r\n\r\n\t\t\t\tfooter * {\r\n\t\t\t\t\tcolor: white;\r\n\t\t\t\t}\r\n\r\n\t\t\t\t#social-links img {\r\n\t\t\t\t\twidth: 1.5rem;\r\n\t\t\t\t\theight: 1.5rem;\r\n\t\t\t\t}\r\n            </style><header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = newNavComponent().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = content.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<footer><p>&copy; Herbal Bones ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", ui.GetYear()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/components/layout.templ`, Line: 60, Col: 60}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</p><div id=\"social-links\"><a href=\"https://www.instagram.com/herbal_bones\" target=\"_blank\"><img src=\"/static/icons/instagram.svg\"></a></div></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func navComponent() templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<style>\r\n\t\tnav {\r\n\t\t\tposition: absolute;\r\n\t\t\ttop: 0;\r\n\t\t\tz-index: 1000;\r\n\t\t\twidth: 100%;\r\n\t\t\theight: 8rem;\r\n\t\t\tbackground: rgb(0,0,0);\r\n\t\t\tbackground: linear-gradient(0deg, rgba(0,0,0,0) 0%, rgba(0,0,0,0.7) 50%, rgba(0,0,0,0.95) 100%);\r\n\t\t}\r\n\r\n\t\t#nav-content-container {\r\n\t\t\tdisplay: flex;\r\n\t\t\tflex-direction: row;\r\n\t\t\tjustify-content: center;\r\n\t\t\talign-items: center;\r\n\t\t\theight: 6rem;\r\n\t\t}\r\n\r\n\t\tnav h1 {\r\n\t\t\tcolor: white;\r\n\t\t\twidth: 33%;\r\n\t\t\tpadding-left: 3rem;\r\n\t\t\tfont-size: 40px;\r\n\t\t}\r\n\r\n\t\tnav ul {\r\n\t\t\tlist-style: none;\r\n\t\t\tdisplay: flex;\r\n\t\t\tflex-direction: row;\r\n\t\t\tjustify-content: center;\r\n\t\t\tgap: 1.5rem;\r\n\t\t\twidth: min-content;\r\n\t\t\tmargin-inline: auto;\r\n\t\t}\r\n\t\tnav a {\r\n\t\t\tcolor: white;\r\n\t\t\ttext-decoration: none;\r\n\t\t}\r\n\r\n\t\t#social-bar {\r\n\t\t\twidth: 33%;\r\n\t\t\tdisplay: flex;\r\n\t\t\tjustify-content: end;\r\n\t\t\tgap: 1rem;\r\n\t\t\tpadding-right: 3rem;\r\n\t\t}\r\n\r\n\t\t#social-bar img {\r\n\t\t\twidth: 2rem;\r\n\t\t\theight: 2rem;\r\n\t\t}\r\n\r\n\t\t@media only screen and (max-width: 650px) {\r\n\t\t\tnav {\r\n\t\t\t\tdisplay: none;\r\n\t\t\t}\r\n\t\t}\r\n\t\t\t\t\r\n\t</style><nav><div id=\"nav-content-container\"><h1>Herbal Bones</h1><ul><li><a href=\"/\">Home</a></li><li><a href=\"/shop\">Shop</a></li><li><a href=\"/contact\">Contact</a></li></ul><div id=\"social-bar\"><a href=\"https://www.instagram.com/herbal_bones\" target=\"_blank\"><img src=\"/static/icons/instagram.svg\"></a></div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func newNavComponent() templ.Component {
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
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<style>\r\n\t\t#outer-nav {\r\n\t\t\tposition: absolute;\r\n\t\t\ttop: 0;\r\n\t\t\tz-index: 1000;\r\n\t\t\twidth: 100%;\r\n\t\t\theight: 8rem;\r\n\t\t\tbackground: rgb(0,0,0);\r\n\t\t\tbackground: linear-gradient(0deg, rgba(0,0,0,0) 0%, rgba(0,0,0,0.7) 50%, rgba(0,0,0,0.95) 100%);\r\n\t\t}\r\n\r\n\t\t#nav-container {\r\n\t\t\tdisplay: flex;\r\n\t\t\tflex-direction: row;\r\n\t\t\talign-items: center;\r\n\t\t\theight: 6rem;\r\n\t\t}\r\n\r\n\t\t#nav-container h1 {\r\n\t\t\tmargin-top: 2rem;\r\n\t\t\tcolor: white;\r\n\t\t\tpadding-left: 3rem;\r\n\t\t\tfont-size: 40px;\r\n\t\t}\r\n\r\n\t\tnav a {\r\n\t\t\tcolor: white;\r\n\t\t\ttext-decoration: none;\r\n\t\t\tdisplay: block;\r\n\t\t\tmargin-block: 1rem;\r\n\t\t}\r\n\r\n\t\t#pop-out {\r\n\t\t\tposition: absolute;\r\n\t\t\ttop: 0;\r\n\t\t\tleft: 0;\r\n\t\t\theight: 100vh;\r\n\t\t\twidth: 100%;\r\n\t\t\tbackground-color: rgba(0,0,0,0.5);\r\n\t\t\tdisplay: none;\r\n\t\t}\r\n\r\n\t\t#pop-out-aside {\r\n\t\t\tbox-sizing: border-box;\r\n\t\t\theight: 100%;\r\n\t\t\twidth: 20rem;\r\n\t\t\tbackground-color: #38383b; \r\n\t\t\tpadding: 3rem;\r\n\t\t\tdisplay: flex;\r\n\t\t\tflex-direction: column;\r\n\t\t\tjustify-content: space-between;\r\n\t\t}\r\n\r\n\t\t#pop-out-aside h1 {\r\n\t\t\tmargin-top: 0;\r\n\t\t\tpadding-left: 0;\r\n\t\t}\r\n\r\n\t\t#hamburger {\r\n\t\t\twidth: 2rem;\r\n\t\t\theight: 2rem;\t\r\n\t\t\tmargin-left: 3rem;\r\n\t\t}\r\n\r\n\t\t#hamburger:hover {\r\n\t\t\tcursor: pointer;\r\n\t\t}\r\n\r\n\r\n\t\t@media only screen and (min-width: 650px) {\r\n\t\t\t#hamburger {\r\n\t\t\t\tdisplay: none;\r\n\t\t\t}\r\n\r\n\t\t\t#nav-container h1 {\r\n\t\t\t\twidth: 100%;\r\n\t\t\t}\r\n\r\n\t\t\t#pop-out {\r\n\t\t\t\tposition: relative;\r\n\t\t\t\tdisplay: block !important;\r\n\t\t\t\theight: 100%;\r\n\t\t\t\tbackground: none;\r\n\t\t\t}\r\n\r\n\t\t\t#pop-out-aside nav {\r\n\t\t\t\twidth: 250px\r\n\t\t\t}\r\n\r\n\t\t\t#pop-out-aside {\r\n\t\t\t\tdisplay: flex;\r\n\t\t\t\tflex-direction: row;\r\n\t\t\t\theight: 100%;\r\n\t\t\t\tpadding-left: 0;\r\n\t\t\t\twidth: calc(100vw / 2 + 125px);\r\n\t\t\t\tbackground: none;\r\n\t\t\t}\r\n\r\n\t\t\t#pop-out-aside h1 {\r\n\t\t\t\tdisplay: none;\r\n\t\t\t}\r\n\r\n\t\t\tnav a {\r\n\t\t\t\tdisplay: inline;\r\n\t\t\t\tmargin-inline: 1rem;\r\n\t\t\t}\r\n\t\t}\r\n\t</style><div id=\"outer-nav\"><div id=\"nav-container\"><div id=\"hamburger\"><img src=\"/static/icons/bars.svg\"></div><h1>Herbal Bones</h1><div id=\"pop-out\"><aside id=\"pop-out-aside\"><div><h1>Herbal Bones</h1><nav><a href=\"/\">Home</a> <a href=\"/shop\">Shop</a> <a href=\"/contact\">Contact</a></nav></div><div id=\"social-links\"><a href=\"https://www.instagram.com/herbal_bones\" target=\"_blank\"><img src=\"/static/icons/instagram.svg\"></a></div></aside></div></div></div><script>\r\n\t\tfunction onHamburgerClick() {\r\n\t\t\tvar popOut = document.getElementById(\"pop-out\");\r\n\t\t\tpopOut.style.display = \"block\";\r\n\t\t}\r\n\r\n\t\tfunction onPopoutClick() {\r\n\t\t\tvar popOut = document.getElementById(\"pop-out\");\r\n\t\t\tpopOut.style.display = \"none\";\r\n\t\t}\r\n\r\n\t\tfunction onPopoutAsideClick(e) {\r\n\t\t\te.stopPropagation()\r\n\t\t}\r\n\r\n\t\t(() => {\r\n\t\t\tdocument.getElementById(\"hamburger\").addEventListener(\"click\", onHamburgerClick);\r\n\t\t\tdocument.getElementById(\"pop-out\").addEventListener(\"click\", onPopoutClick);\r\n\t\t\tdocument.getElementById(\"pop-out-aside\").addEventListener(\"click\", onPopoutAsideClick);\r\n\t\t})()\r\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
