package server

import (
	"github.com/goadesign/goa"
	"github.com/oosidat/go-url-shortener/app"
	"github.com/oosidat/go-url-shortener/stores"
)

// ShortURLController implements the short_url resource.
type ShortURLController struct {
	*goa.Controller
	store stores.IStorage
}

// NewShortURLController creates a short_url controller.
func NewShortURLController(service *goa.Service, storage stores.IStorage) *ShortURLController {
	return &ShortURLController{
		Controller: service.NewController("ShortURLController"),
		store: storage,
	}
}

// CreateShortURL runs the create_short_url action.
func (c *ShortURLController) CreateShortURL(ctx *app.CreateShortURLShortURLContext) error {
	longURL := ctx.Payload.URL
	shortURL := c.store.Save(longURL)
	res := &app.GoaExampleShortURL{LongURL: &longURL, ShortURL: shortURL}
	return ctx.Created(res)
}

// GetShortURL runs the get_short_url action.
func (c *ShortURLController) GetShortURL(ctx *app.GetShortURLShortURLContext) error {
	shortURL := ctx.ShortURLHash
	longURL, err := c.store.Load(shortURL)

	if err != nil {
		return ctx.NotFound()
	}

	res := &app.GoaExampleShortURL{LongURL: &longURL, ShortURL: shortURL}
	return ctx.OK(res)
}

// RedirectShortURL runs the redirect_short_url action.
func (c *ShortURLController) RedirectShortURL(ctx *app.RedirectShortURLShortURLContext) error {
	shortURL := ctx.ShortURLHash
	longURL, err := c.store.Load(shortURL)

	if err != nil {
		return ctx.NotFound()
	}

	header := ctx.ResponseWriter.Header()
	header.Set("Location", longURL)
	return ctx.MovedPermanently()
}
