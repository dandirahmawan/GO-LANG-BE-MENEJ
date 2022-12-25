package services

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/input"
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/dandirahmawan/menej_api_go/response_base"
	"github.com/gin-gonic/gin"
)

type ModelResp []model.LabelsModel

func GetLabelsByProjectId(ctx *gin.Context) ModelResp {
	var id string = ctx.Param("id")
	data := model.FindLabelsByProjectId(id)
	return data
}

func SaveDataLabel(ctx *gin.Context) interface{} {
	var par model.LabelsModel
	ctx.Bind(&par)
	par.Date = time.Now()
	par.Save()
	return par
}

func DeleteLabel(input input.LabelDelete) interface{} {
	isExist := IsExistsLabel(input.ProjectId, input.Label)
	if !isExist {
		return response_base.ResponseMsg(false, "Data label not found")
	} else {
		var dataLabel model.LabelsModel
		dataLabel.ProjectId = input.ProjectId
		dataLabel.Label = input.Label
		dataLabel.Delete()
		return response_base.ResponseMsg(true, "Delete label successfully")
	}
}

func IsExistsLabel(projectId string, label string) bool {
	data := model.FindLabelByIdAndLabel(projectId, label)
	return len(data) > 0
}
