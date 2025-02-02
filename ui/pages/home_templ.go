// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "web.herbalbones.com/ui/components"

func Home() templ.Component {
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
		templ_7745c5c3_Err = components.Layout(homeContent()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func homeContent() templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<style>\r\n        body {\r\n            margin: 0;\r\n        }\r\n\r\n        #contact-section {\r\n            height: 45rem;\r\n            background-image: url(\"/static/images/contact-background.jpg\");\r\n            background-position: bottom;\r\n            background-repeat: no-repeat;\r\n            background-size: cover;\r\n        }\r\n    </style><main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = heroSection().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = shopSection().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = aboutSection().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<section id=\"contact-section\"></section></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func heroSection() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<style>\r\n        #hero-section {\r\n            aspect-ratio: 1.9;\r\n            background-image: url(\"/static/images/hero-background.jpg\");\r\n            background-position: center;\r\n            background-repeat: no-repeat;\r\n            background-size: cover;\r\n        }\r\n            #hero-section > div {\r\n                height: 100%;\r\n                width: 100%;\r\n                background-color: rgba(0, 0, 0, 0.3);\r\n                display: flex;\r\n                flex-direction: column;\r\n                align-items: center;\r\n                justify-content: center;\r\n            }\r\n            #hero-section h2 {\r\n                text-align: center;\r\n                color: white;\r\n                font-size: 5em;\r\n            }\r\n    </style><section id=\"hero-section\"><div><h2>bring fantasy to the real world</h2></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func shopSection() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<style>\r\n        #shop-section {\r\n            background-color: #a8a6a1;\r\n        }\r\n\r\n            #shop-section h3 {\r\n                text-align: center;\r\n                margin-block: 0;\r\n                padding-block: 3rem;\r\n                font-size: 2.25em;\r\n                color: white;\r\n            }\r\n            #artwork-row {\r\n                display: grid;\r\n                grid-template-columns: 1fr;\r\n                gap: 2rem;\r\n                margin-inline: 2rem;\r\n                max-width: 1024px;\r\n            }\r\n\r\n            #artwork-row .artwork-card {\r\n                max-width: 300px;\r\n                margin-inline: auto;\r\n            }\r\n\r\n            @media only screen and (min-width: 600px) {\r\n                #artwork-row {\r\n                    grid-template-columns: repeat(3, 1fr);\r\n                }\r\n                #artwork-row .artwork-card {\r\n                    max-width: 100%;\r\n                    margin-inline: 0;\r\n                }\r\n            }\r\n\r\n            @media only screen and (min-width: 1024px) {\r\n\r\n                #artwork-row {\r\n                    margin-inline: auto;\r\n                }\r\n\r\n            }\r\n\r\n\r\n        #action-row {\r\n            display: flex;\r\n            justify-content: center;\r\n            padding-block: 2rem;\r\n        }\r\n\r\n        #action-row > a {\r\n            padding: 1rem;\r\n            padding-inline: 1.5rem;\r\n            background-color: white;\r\n            color: #a5a6a7;\r\n            text-decoration: none;\r\n            border-radius: 3px;\r\n        }\r\n    </style><section id=\"shop-section\"><h3>fantasy artwork</h3><div id=\"artwork-row\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = artworkCard("/static/images/artwork-everyday.jpg", "red necklace and earrings", "everyday wear").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = artworkCard("/static/images/artwork-statements.png", "red necklace", "bigger statements").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = artworkCard("/static/images/artwork-fantasy.jpeg", "red necklace with pearls", "fantasy").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</div><div id=\"action-row\"><a href=\"/shop\">SHOP</a></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func aboutSection() templ.Component {
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
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<style>\r\n        #about-section {\r\n            background-color: #dfe0e1;\r\n            padding-block: 3rem;\r\n        }\r\n        #about-section img {\r\n            max-width: 500px;\r\n            aspect-ratio: 0.9;\r\n            object-fit: cover;\r\n        }\r\n        #about-section p {\r\n            font-size: 1.5em;\r\n            text-align: center;\r\n            margin-inline: 1.25rem;\r\n            color: #303030;\r\n        }\r\n\r\n        #about-section .content-wrapper {\r\n            display: flex;\r\n            flex-direction: column;\r\n            align-items: center;\r\n            margin-left: auto;\r\n        }\r\n\r\n        @media only screen and (min-width: 1024px) {\r\n            #about-section .content-wrapper {\r\n                display: grid;\r\n                grid-template-columns: 1fr 500px;\r\n                align-items: center;\r\n                margin-left: auto;\r\n            }\r\n            #about-section img {\r\n                width: 100%;\r\n            }\r\n        }\r\n\r\n    </style><section id=\"about-section\"><div class=\"content-wrapper\"><p>Discover beautiful, handcrafted jewelry that brings harmony and balance to your soul. Each piece is carefully made to blend fantasy with reality, offering a meaningful gift for loved ones or a way to uplift your own spirit. Created in 2023, this jewelry reflects my passion and dedication to perfecting my skills and appreciating the beauty of jewelry. Dive into nature-inspired jewelry to elevate your cosplay experience, bringing an otherworldly charm to your characters and stories.</p><img src=\"/static/images/about.jpeg\" alt=\"girl posing wearing red necklace\"></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func artworkCard(imgUrl string, imgAlt string, text string) templ.Component {
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
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<style>\r\n        .artwork-card {\r\n            width: 100%;\r\n            height: 100%;\r\n            display: flex;\r\n            flex-direction: column;\r\n            align-items: center;\r\n        }\r\n        \r\n            .artwork-card img {\r\n                width: 100%;\r\n                height: 100%;\r\n                aspect-ratio: 0.7;\r\n                object-fit: cover;\r\n            }\r\n\r\n            .artwork-card p {\r\n                color: white;\r\n                font-size: 1.75em;\r\n            }\r\n    </style><div class=\"artwork-card\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(imgUrl)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/pages/home.templ`, Line: 211, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\" alt=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(imgAlt)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/pages/home.templ`, Line: 211, Col: 34}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"><p class=\"serif\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(text)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/pages/home.templ`, Line: 212, Col: 25}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
