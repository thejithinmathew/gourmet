package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thejithinmathew/gourmet/pkg/models"
	"github.com/thejithinmathew/gourmet/pkg/utils"
)

func (c *Controller) PatchBook(ctx *gin.Context) {
	req := &models.BookReq{}
	parseErr := utils.GetRequest(ctx, req)
	if parseErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, parseErr)
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
