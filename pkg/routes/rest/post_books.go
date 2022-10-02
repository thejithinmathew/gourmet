package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thejithinmathew/gourmet/pkg/models"
	"github.com/thejithinmathew/gourmet/pkg/models/errors"
)

func (c *Controller) PostBooks(ctx *gin.Context) {
	var books models.Books
	err := ctx.ShouldBindJSON(&books)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, err.Error()))
	}
}
