package main

import (
	"context"
	"log"

	"github.com/thejithinmathew/gourmet/pkg/clients"
	"github.com/thejithinmathew/gourmet/pkg/config"
	"github.com/thejithinmathew/gourmet/pkg/routes"
)

func main() {
	config.Init()
	cfg := config.Get()
	Client, err := clients.Init(&cfg)
	if err != nil {
		log.Fatal("unable to initialize clients", err)
	}
	defer func() {
		if err = Client.DBClient.Disconnect(context.Background()); err != nil {
			log.Fatal("unable to initialize clients", err)
		}
	}()
	routes.New(&cfg, Client)
}
