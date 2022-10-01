package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thejithinmathew/gourmet/pkg/clients"
	"github.com/thejithinmathew/gourmet/pkg/config"
	"github.com/thejithinmathew/gourmet/pkg/routes/middlewares"
	"github.com/thejithinmathew/gourmet/pkg/routes/rest"
)

func New(cfg *config.Config, clients *clients.Clients) {
	fmt.Println("checking routes")
	engine := gin.New()
	route := engine.Group("")
	route.Use(
		middlewares.PanicMiddleware(),
	)
	c := &rest.Controller{
		Clients: clients,
		Config:  cfg,
	}
	route.POST("/book", c.PostBook)
	route.PATCH("/book", c.PatchBook)
	engine.Run()
}
