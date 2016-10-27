package design

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

var ShortUrlMedia = dsl.MediaType("application/vnd.goa.example.short_url+json", func() {
	dsl.Description("Media for Short url")
	dsl.Attributes(func() {
		dsl.Attribute("shortUrl", goa.String)
		dsl.Required("shortUrl")
		dsl.Attribute("longUrl", goa.String)
	})
	dsl.View("default", func() {
		dsl.Attribute("shortUrl")
		dsl.Required("shortUrl")
		dsl.Attribute("longUrl")
	})
})
