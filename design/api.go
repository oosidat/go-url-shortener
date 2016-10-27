package design

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

func init() {

	dsl.API("url-shortener", func() {
		dsl.Title("URL Shortener")
		dsl.Scheme("http")
		dsl.Version("1.0")
		dsl.BasePath("/links")
		dsl.Produces("application/json")
		dsl.Consumes("application/json")
	})

	dsl.Resource("short_url", func() {
		dsl.Action("get_short_url", func() {
			dsl.Description("Get a short url by hash")
			dsl.Routing(dsl.GET("/decode/:short_url_hash"))
			dsl.Response(goa.OK, func() {
				dsl.Media(ShortUrlMedia)
			})
			dsl.Response(goa.BadRequest)
			dsl.Response(goa.NotFound)
		})

		dsl.Action("redirect_short_url", func() {
			dsl.Description("Redirect a short url")
			dsl.Routing(dsl.GET("/redirect/:short_url_hash"))
			dsl.Response(goa.MovedPermanently, func() {
				dsl.Headers(func() {
					dsl.Header("Location", goa.String, func() {
						dsl.Description("The URI to the provided short URL")
					})
				})
			})
			dsl.Response(goa.BadRequest)
			dsl.Response(goa.NotFound)
		})

		dsl.Action("create_short_url", func() {
			dsl.Description("Create a short url")
			dsl.Routing(dsl.POST("/"))
			dsl.Payload(ShortUrlCreatePayload)
			dsl.Response(goa.Created, ShortUrlMedia)
			dsl.Response(goa.BadRequest)
			dsl.Response(goa.InternalServerError)
		})

	})
}
