package main

import (
	"log"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
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

	cfg := stores.Config{
		Cluster: []string{"127.0.0.1"},
		Keyspace: "example",
	}
	store := &stores.Cassandra{}
	err := store.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	c := server.NewShortURLController(service, store)
	app.MountShortURLController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
