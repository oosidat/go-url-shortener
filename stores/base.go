package stores

import (
	"github.com/oosidat/go-url-shortener/app"
)

type Store interface {
	Code() (string, error)
	Save(string) (string, error)
	SavePayload(*app.ShortURLCreatePayload) (string, error)
	Load(string) (string, error)
	LoadRecord(string) (app.GoaExampleShortURL, error)
}
