//************************************************************************//
// API "url-shortener": Application Controllers
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

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ShortURLController is the controller interface for the ShortURL actions.
type ShortURLController interface {
	goa.Muxer
	CreateShortURL(*CreateShortURLShortURLContext) error
	GetShortURL(*GetShortURLShortURLContext) error
	RedirectShortURL(*RedirectShortURLShortURLContext) error
}

// MountShortURLController "mounts" a ShortURL resource controller on the given service.
func MountShortURLController(service *goa.Service, ctrl ShortURLController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateShortURLShortURLContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ShortURLCreatePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.CreateShortURL(rctx)
	}
	service.Mux.Handle("POST", "/", ctrl.MuxHandler("CreateShortURL", h, unmarshalCreateShortURLShortURLPayload))
	service.LogInfo("mount", "ctrl", "ShortURL", "action", "CreateShortURL", "route", "POST /")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetShortURLShortURLContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.GetShortURL(rctx)
	}
	service.Mux.Handle("GET", "/dec/:short_url_hash", ctrl.MuxHandler("GetShortURL", h, nil))
	service.LogInfo("mount", "ctrl", "ShortURL", "action", "GetShortURL", "route", "GET /dec/:short_url_hash")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRedirectShortURLShortURLContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.RedirectShortURL(rctx)
	}
	service.Mux.Handle("GET", "/red/:short_url_hash", ctrl.MuxHandler("RedirectShortURL", h, nil))
	service.LogInfo("mount", "ctrl", "ShortURL", "action", "RedirectShortURL", "route", "GET /red/:short_url_hash")
}

// unmarshalCreateShortURLShortURLPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateShortURLShortURLPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &shortURLCreatePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
