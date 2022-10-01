package utils

import (
	"github.com/gin-gonic/gin"
)

func GetRequest(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindJSON(obj)
	return err
}
