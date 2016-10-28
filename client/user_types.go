//************************************************************************//
// API "url-shortener": Application User Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/oosidat/go-url-shortener/design
// --out=$(GOPATH)/src/github.com/oosidat/go-url-shortener
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import "github.com/goadesign/goa"

// Payload for creating shortened url
type shortURLCreatePayload struct {
	// The optional context for the url being shortened
	Context map[string]string `form:"context,omitempty" json:"context,omitempty" xml:"context,omitempty"`
	// The url to be shortened
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// Validate validates the shortURLCreatePayload type instance.
func (ut *shortURLCreatePayload) Validate() (err error) {
	if ut.URL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}

	return
}

// Publicize creates ShortURLCreatePayload from shortURLCreatePayload
func (ut *shortURLCreatePayload) Publicize() *ShortURLCreatePayload {
	var pub ShortURLCreatePayload
	if ut.Context != nil {
		pub.Context = ut.Context
	}
	if ut.URL != nil {
		pub.URL = *ut.URL
	}
	return &pub
}

// Payload for creating shortened url
type ShortURLCreatePayload struct {
	// The optional context for the url being shortened
	Context map[string]string `form:"context,omitempty" json:"context,omitempty" xml:"context,omitempty"`
	// The url to be shortened
	URL string `form:"url" json:"url" xml:"url"`
}

// Validate validates the ShortURLCreatePayload type instance.
func (ut *ShortURLCreatePayload) Validate() (err error) {
	if ut.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}

	return
}
