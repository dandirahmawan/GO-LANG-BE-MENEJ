package services

import (
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

type ModelResp []model.LabelsModel

func GetLabelsByProjectId(ctx *gin.Context) ModelResp {
	var id string = ctx.Param("id")
	data := model.FindLabelsByProjectId(id)
	return data
}
