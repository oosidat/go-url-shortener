//************************************************************************//
// API "url-shortener": Application Contexts
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
)

// CreateShortURLShortURLContext provides the short_url create_short_url action context.
type CreateShortURLShortURLContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *ShortURLCreatePayload
}

// NewCreateShortURLShortURLContext parses the incoming request URL and body, performs validations and creates the
// context used by the short_url controller create_short_url action.
func NewCreateShortURLShortURLContext(ctx context.Context, service *goa.Service) (*CreateShortURLShortURLContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateShortURLShortURLContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateShortURLShortURLContext) Created(r *GoaExampleShortURL) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.short_url+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateShortURLShortURLContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateShortURLShortURLContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// GetShortURLShortURLContext provides the short_url get_short_url action context.
type GetShortURLShortURLContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ShortURLHash string
}

// NewGetShortURLShortURLContext parses the incoming request URL and body, performs validations and creates the
// context used by the short_url controller get_short_url action.
func NewGetShortURLShortURLContext(ctx context.Context, service *goa.Service) (*GetShortURLShortURLContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetShortURLShortURLContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramShortURLHash := req.Params["short_url_hash"]
	if len(paramShortURLHash) > 0 {
		rawShortURLHash := paramShortURLHash[0]
		rctx.ShortURLHash = rawShortURLHash
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetShortURLShortURLContext) OK(r *GoaExampleShortURL) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.short_url+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GetShortURLShortURLContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetShortURLShortURLContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// RedirectShortURLShortURLContext provides the short_url redirect_short_url action context.
type RedirectShortURLShortURLContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ShortURLHash string
}

// NewRedirectShortURLShortURLContext parses the incoming request URL and body, performs validations and creates the
// context used by the short_url controller redirect_short_url action.
func NewRedirectShortURLShortURLContext(ctx context.Context, service *goa.Service) (*RedirectShortURLShortURLContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RedirectShortURLShortURLContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramShortURLHash := req.Params["short_url_hash"]
	if len(paramShortURLHash) > 0 {
		rawShortURLHash := paramShortURLHash[0]
		rctx.ShortURLHash = rawShortURLHash
	}
	return &rctx, err
}

// MovedPermanently sends a HTTP response with status code 301.
func (ctx *RedirectShortURLShortURLContext) MovedPermanently() error {
	ctx.ResponseData.WriteHeader(301)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *RedirectShortURLShortURLContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RedirectShortURLShortURLContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
