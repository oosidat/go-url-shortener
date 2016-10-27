package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateShortURLShortURLPath computes a request path to the create_short_url action of short_url.
func CreateShortURLShortURLPath() string {
	return fmt.Sprintf("/links")
}

// Create a short url
func (c *Client) CreateShortURLShortURL(ctx context.Context, path string, payload *ShortURLCreatePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateShortURLShortURLRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateShortURLShortURLRequest create the request corresponding to the create_short_url action endpoint of the short_url resource.
func (c *Client) NewCreateShortURLShortURLRequest(ctx context.Context, path string, payload *ShortURLCreatePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// GetShortURLShortURLPath computes a request path to the get_short_url action of short_url.
func GetShortURLShortURLPath(shortURLHash string) string {
	return fmt.Sprintf("/links/decode/%v", shortURLHash)
}

// Get a short url by hash
func (c *Client) GetShortURLShortURL(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetShortURLShortURLRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetShortURLShortURLRequest create the request corresponding to the get_short_url action endpoint of the short_url resource.
func (c *Client) NewGetShortURLShortURLRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RedirectShortURLShortURLPath computes a request path to the redirect_short_url action of short_url.
func RedirectShortURLShortURLPath(shortURLHash string) string {
	return fmt.Sprintf("/links/redirect/%v", shortURLHash)
}

// Redirect a short url
func (c *Client) RedirectShortURLShortURL(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRedirectShortURLShortURLRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRedirectShortURLShortURLRequest create the request corresponding to the redirect_short_url action endpoint of the short_url resource.
func (c *Client) NewRedirectShortURLShortURLRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
