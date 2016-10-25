package main

import (
	"log"
	"path/filepath"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/mitchellh/go-homedir"
	"github.com/oosidat/go-url-shortener/app"
	"github.com/oosidat/go-url-shortener/stores"
	"github.com/oosidat/go-url-shortener/server"
)

func main() {
	// Create service
	service := goa.New("url-shortener")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "short_url" controller
	dir, _ := homedir.Dir()
	store := &stores.Filesystem{}
	err := store.Init(filepath.Join(dir, "go-url-shortener"))
	if err != nil {
		log.Fatal(err)
	}

	c := server.NewShortURLController(service, store)
	app.MountShortURLController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
