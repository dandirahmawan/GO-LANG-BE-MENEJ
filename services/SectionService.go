package services

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

func GetDataSection(ctx *gin.Context) interface{} {
	pi := ctx.Param("id")
	data := model.FindSectionOnlyByProjectId(pi)
	return data
}

func SaveSection(param model.SectionModel) interface{} {
	param.CreatedDate = time.Now()
	param.Id = commons.GeneratdUUID(20)
	param.Save()

	var resp struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}

	data := model.FindSectionOnlyByProjectId(param.ProjectId)
	resp.Data = data
	resp.Success = true

	return resp
}

func DeleteSection(id string, userid string) interface{} {
	var data int16 = model.DeleteSectionById(id)

	type rsp struct {
		Succcess bool   `json:"success"`
		Message  string `json:"message"`
	}

	var resp rsp
	if data == 0 {
		resp.Succcess = false
		resp.Message = "Something wrong when delete data"
		return resp
	}

	resp.Succcess = true
	resp.Message = "Delete section successfully"
	return resp
}
