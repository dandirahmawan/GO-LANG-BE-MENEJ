package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/input"
	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var input input.InputUser
	ctx.BindJSON(&input)

	resp := services.RegisterUser(input)
	ctx.IndentedJSON(http.StatusOK, resp)
}
