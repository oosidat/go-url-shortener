package design

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

var ShortUrlCreatePayload = dsl.Type("short_url_create_payload", func() {
	dsl.Description("Payload for creating shortened url")
	dsl.Attribute("url", goa.String)
	dsl.Required("url")
})
