package internal

import "github.com/gin-gonic/gin"

func IsPing() gin.HandlerFunc {
	return func(context *gin.Context) {
		value := context.Request.Header.Get("X-PING")

		if value == "ping" {
			context.Writer.Header().Set("X-PONG", "pong")
		}

		context.Next()
	}
}
