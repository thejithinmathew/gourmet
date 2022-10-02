package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thejithinmathew/gourmet/pkg/clients"
	"github.com/thejithinmathew/gourmet/pkg/config"
	"github.com/thejithinmathew/gourmet/pkg/models"
	"github.com/thejithinmathew/gourmet/pkg/models/errors"
	"github.com/thejithinmathew/gourmet/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	Config  *config.Config
	Clients *clients.Clients
}

func (c *Controller) PostBook(ctx *gin.Context) {
	logger := utils.GetLogger(ctx)
	logger.With().Str("METHOD", "PostBook")
	req := &models.BookReq{}
	book := models.Book{}
	parseErr := utils.GetRequest(ctx, req)
	if parseErr != nil {
		logger.Error().Err(parseErr).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, parseErr.Error()))
		return
	}
	if req.Name == "" {
		req.Name = "random name"
	} else {
		if er := utils.ValidateName(req.Name); er != nil {
			logger.Error().Err(er).Send()
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, er.Error()))
			return
		}
	}
	res := c.Clients.DBClient.Database("dbset").Collection("books").FindOne(ctx, bson.M{"name": req.Name})
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		logger.Err(res.Err())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.NewError(http.StatusInternalServerError, res.Err().Error()))
		return
	}
	if res.Err() == nil {
		_ = res.Decode(&book)
		ctx.JSON(http.StatusOK, book)
		return
	}
	book = models.Book{
		Name:       req.Name,
		Generation: 1,
		Author:     req.Author,
		ISBN:       req.ISBN,
	}
	_, err := c.Clients.DBClient.Database("dbset2").Collection("books").InsertOne(ctx, book)
	if err != nil {
		logger.Error().Str("ERROR", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, &book)
}
