package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
			return file + ":" + strconv.Itoa(line)
		}

		reqID := ctx.Request.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		} else {
			if _, err := uuid.Parse(reqID); err != nil {
				reqID = uuid.New().String()
			}
		}

		subLogger := log.With().Str("REQ_ID", reqID).Logger()

		subLogger.Info().
			Str("path", ctx.Request.URL.Path).
			Str("REQ_ID", reqID).
			Str("method", ctx.Request.Method).
			Msg("Incoming Request")

		ctx.Set("X-Request-ID", reqID)
		ctx.Request.Header.Set("X-Request-ID", reqID)
		ctx.Writer.Header().Set("X-Request-ID", reqID)
		ctx.Set("logger", subLogger)
		ctx.Next()

		subLogger.Info().
			Int("status", ctx.Writer.Status()).
			Msg("Outgoing Response")
	}
}
