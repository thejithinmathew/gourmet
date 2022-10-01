package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thejithinmathew/gourmet/pkg/models/errors"
)

func PanicMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := recover(); err != nil {
			newErr := errors.NewError(http.StatusInternalServerError, "panic caught by middleware")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, newErr)
			return
		}
	}
}
