package services

import (
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

func GetDataSection(ctx *gin.Context) interface{} {
	pi := ctx.Param("id")
	return model.FindSectionOnlyByProjectId(pi)
}
