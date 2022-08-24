package services

import (
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

func GetProjectTeamByProjectId(ctx *gin.Context) interface{} {
	var projectID string = ctx.Param("id")
	data := model.FindViewProjectTeamByProjectId(projectID)
	return data
}
