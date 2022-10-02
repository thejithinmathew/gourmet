package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thejithinmathew/gourmet/pkg/models"
	"github.com/thejithinmathew/gourmet/pkg/models/errors"
	"github.com/thejithinmathew/gourmet/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) PatchBook(ctx *gin.Context) {
	logger := utils.GetLogger(ctx)
	req := &models.UpdateBookReq{}
	parseErr := utils.GetRequest(ctx, req)
	id := ctx.Param("id")
	if parseErr != nil {
		logger.Error().Str("ERROR", parseErr.Error()).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, parseErr.Error()))
		return
	}

	updateRes, err := c.Clients.DBClient.Database("dbset").Collection("books").UpdateOne(ctx,
		bson.M{
			"id": id,
		},
		bson.M{
			"$set": bson.M{
				"name": req.Name,
			},
		},
	)
	if err != nil {
		logger.Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	if updateRes.MatchedCount == 0 {
		err := fmt.Errorf("no documents matched")
		logger.Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusNotFound, errors.NewError(http.StatusNotFound, "document not found"))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
