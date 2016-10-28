//************************************************************************//
// API "url-shortener": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/oosidat/go-url-shortener/design
// --out=$(GOPATH)/src/github.com/oosidat/go-url-shortener
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// Media for Short url (default view)
//
// Identifier: application/vnd.goa.example.short_url+json; view=default
type GoaExampleShortURL struct {
	Context  map[string]string `form:"context,omitempty" json:"context,omitempty" xml:"context,omitempty"`
	LongURL  *string           `form:"longUrl,omitempty" json:"longUrl,omitempty" xml:"longUrl,omitempty"`
	ShortURL string            `form:"shortUrl" json:"shortUrl" xml:"shortUrl"`
}

// Validate validates the GoaExampleShortURL media type instance.
func (mt *GoaExampleShortURL) Validate() (err error) {
	if mt.ShortURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "shortUrl"))
	}

	return
}
