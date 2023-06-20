package middlewares

import "github.com/gin-gonic/gin"

func Set(key string, value any) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(key, value)
		ctx.Next()
	}
}
