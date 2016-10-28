package design

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

var ShortUrlMedia = dsl.MediaType("application/vnd.goa.example.short_url+json", func() {
	dsl.Description("Media for Short url")
	dsl.Attributes(func() {
		dsl.Attribute("shortUrl", goa.String, func() {
			dsl.Example("8iLV0Z02q")
		})
		dsl.Required("shortUrl")
		dsl.Attribute("longUrl", goa.String, func() {
			dsl.Example("https://github.com/oosidat/go-url-shortener")
		})
		dsl.Attribute("context", dsl.HashOf(goa.String, goa.String), func() {
			dsl.Example(map[string]string{
				"foo":  "bar",
				"fizz": "buzz",
			})
		})
	})
	dsl.View("default", func() {
		dsl.Attribute("shortUrl")
		dsl.Required("shortUrl")
		dsl.Attribute("longUrl")
		dsl.Attribute("context")
	})
})
