// Package storages allows multiple implementation on how to store short URLs.
package stores

type IStorage interface {
	Code() string
	Save(string) (string, error)
	Load(string) (string, error)
}
