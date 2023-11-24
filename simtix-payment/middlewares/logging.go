package middlewares

import (
	"fmt"
	"simtix/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Starting time request
		startTime := time.Now()

		// Processing request
		ctx.Next()

		// End Time request
		endTime := time.Now()

		// execution time
		latencyTime := endTime.Sub(startTime)

		// Request method
		reqMethod := ctx.Request.Method

		// Request route
		reqUri := ctx.Request.RequestURI

		// status code
		statusCode := ctx.Writer.Status()

		// Request IP
		clientIP := ctx.ClientIP()

		logger.Log.Info(
			fmt.Sprintf(
				"[%s | %d] \"%s\"  |   %s  |    %s", reqMethod, statusCode, reqUri, latencyTime,
				clientIP,
			),
		)

		ctx.Next()
	}
}
