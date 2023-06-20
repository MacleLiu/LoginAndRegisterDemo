package routers

import (
	"LoginAndRegisterDemo/main/handlers"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		// 注册
		userRouter.GET("/register", handlers.UserHandler{}.Register)
		userRouter.POST("/doRegister", handlers.UserHandler{}.DoRegister)

		// 登录
		userRouter.GET("/login", handlers.UserHandler{}.Login)
		userRouter.POST("/doLogin", handlers.UserHandler{}.DoLogin)

		// 忘记密码，修改密码
		userRouter.GET("/forgetPwd", handlers.UserHandler{}.ChangePwd)
		userRouter.POST("/doChangePwd", handlers.UserHandler{}.DoChangePwd)
	}
}
