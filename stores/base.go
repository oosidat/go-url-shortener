// Package storages allows multiple implementation on how to store short URLs.
package stores

type Store interface {
	Code() (string, error)
	Save(string) (string, error)
	Load(string) (string, error)
}
