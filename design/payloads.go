package design

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

var ShortUrlCreatePayload = dsl.Type("short_url_create_payload", func() {
	dsl.Description("Payload for creating shortened url")
	dsl.Attribute("url", goa.String, func() {
		dsl.Description("The url to be shortened")
		dsl.Example("https://github.com/oosidat/go-url-shortener")
	})
	dsl.Required("url")
	dsl.Attribute("context", dsl.HashOf(goa.String, goa.String), func() {
		dsl.Description("The optional context for the url being shortened")
		dsl.Example(map[string]string{
			"foo":  "bar",
			"fizz": "buzz",
		})
	})
})
