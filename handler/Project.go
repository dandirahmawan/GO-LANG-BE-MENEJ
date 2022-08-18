package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

func GetAllProjects(ctx *gin.Context) {
	var resp interface{}
	resp = services.GetAllViewProjects(ctx)
	ctx.IndentedJSON(http.StatusOK, resp)
}
