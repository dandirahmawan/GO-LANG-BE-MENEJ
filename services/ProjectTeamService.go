package services

import (
	"strconv"

	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

func GetProjectTeamByProjectId(ctx *gin.Context) interface{} {
	var projectID string = ctx.Param("id")
	pi, _ := strconv.Atoi(projectID)
	data := model.FindViewProjectTeamByProjectId(int64(pi))
	return data
}
