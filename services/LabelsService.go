package services

import (
	"strconv"

	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

type ModelResp []model.LabelsModel

func GetLabelsByProjectId(ctx *gin.Context) ModelResp {
	var id string = ctx.Param("id")
	ids, _ := strconv.Atoi(id)

	data := model.FindLabelsByProjectId(int64(ids))
	return data
}
