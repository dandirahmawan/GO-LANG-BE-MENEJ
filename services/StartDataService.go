package services

import (
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

func GetStartData(ctx *gin.Context) model.StartData {
	sessionId := ctx.Request.Header.Get("sessionid")
	return model.FindDataStart(sessionId)
}
