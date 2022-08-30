package services

import (
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

func GetDataPermition(ctx *gin.Context) interface{} {
	userId := ctx.Request.Header.Get("userid")
	projectId := ctx.Param("id")
	var data []model.PermitionProject
	data = model.FindDataPermitionByProjectIdAndUserId(userId, projectId)
	return data
}
