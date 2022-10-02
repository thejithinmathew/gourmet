package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/thejithinmathew/gourmet/pkg/models"
	"github.com/thejithinmathew/gourmet/pkg/utils"
)

func (c *Controller) PatchBook(ctx *gin.Context) {
	log.With().Str("METHOD", "PatchBook")
	req := &models.BookReq{}
	parseErr := utils.GetRequest(ctx, req)
	if parseErr != nil {
		log.Error().Str("ERROR", parseErr.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, parseErr)
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
