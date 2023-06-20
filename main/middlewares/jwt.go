package middlewares

import (
	"LoginAndRegisterDemo/main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求携带的 Cookie 信息
		tokenStr, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.String(http.StatusForbidden, "获取cookie错误, 拒绝访问")
			ctx.Abort()
			return
		}
		if tokenStr == "" {
			ctx.String(http.StatusForbidden, "token为空, 拒绝访问")
			ctx.Abort()
			return
		}

		// 解析客户端发送的 token
		token, err := utils.ParseToken(tokenStr)
		if err != nil {
			ctx.String(http.StatusForbidden, "token 解析错误")
			ctx.Abort()
			return
		}

		// 获取 token 中的claims
		claims, ok := token.Claims.(*utils.Claims)
		if !ok || !token.Valid {
			ctx.String(http.StatusForbidden, "token 无效")
			ctx.Abort()
			return
		}

		// 将解析后的有效载荷 claims 写入 gin.Context 中
		ctx.Set("claims", claims)

		ctx.Next()
	}
}
