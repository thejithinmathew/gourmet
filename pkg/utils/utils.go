package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetRequest(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindJSON(obj)
	return err
}

func GetLogger(ctx *gin.Context) zerolog.Logger {
	logging, _ := ctx.Get("logger")
	if logger, isLogger := logging.(zerolog.Logger); isLogger {
		return logger.With().Caller().Logger()
	}
	return log.Logger
}
