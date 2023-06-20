package handlers

import (
	"LoginAndRegisterDemo/main/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultHandler struct{}

func (u DefaultHandler) Index(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		log.Fatal("Get claims err in index")
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"name": claims.(*utils.Claims).UserName,
	})
}
