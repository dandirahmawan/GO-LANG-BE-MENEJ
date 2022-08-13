package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

func GetAllProjects(ctx *gin.Context) {
	var resp interface{}
	resp = model.FindAll()
	ctx.IndentedJSON(http.StatusOK, resp)
}
