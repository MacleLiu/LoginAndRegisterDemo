package routers

import (
	"LoginAndRegisterDemo/main/handlers"
	"LoginAndRegisterDemo/main/middlewares"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {
	defaultRouter := r.Group("/")

	// 使用JWT鉴权
	defaultRouter.Use(middlewares.JWTAuth())
	{
		defaultRouter.GET("/", handlers.DefaultHandler{}.Index)
	}
}
