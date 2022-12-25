package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/input"
	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

func PostLabel(ctx *gin.Context) {
	data := services.SaveDataLabel(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}

func DeleteLabel(ctx *gin.Context) {
	var input input.LabelDelete
	ctx.BindJSON(&input)

	data := services.DeleteLabel(input)
	ctx.IndentedJSON(http.StatusOK, data)
}
