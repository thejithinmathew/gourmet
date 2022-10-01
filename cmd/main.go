package main

import (
	"context"
	"fmt"

	"github.com/thejithinmathew/gourmet/pkg/clients"
	"github.com/thejithinmathew/gourmet/pkg/config"
	"github.com/thejithinmathew/gourmet/pkg/routes"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("Hi")
	config.Init()
	cfg := config.Get()
	Client, err := clients.Init(&cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func() {
		if err = Client.DBClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	f, _ := Client.DBClient.ListDatabaseNames(context.Background(), bson.M{})
	fmt.Println(f)
	fmt.Println(cfg.User + " lol")
	routes.New(&cfg, Client)
}
