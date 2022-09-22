package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

func PostSection(ctx *gin.Context) {
	var par model.SectionModel
	ctx.BindJSON(&par)
	data := services.SaveSection(par)

	ctx.IndentedJSON(http.StatusOK, data)
}

func DeleteSection(ctx *gin.Context) {
	id := ctx.Param("id")
	sessionid := ctx.Request.Header.Get("sessionid")

	dataSession := commons.GetDataSession(sessionid)
	accountId := dataSession.Id

	data := services.DeleteSection(id, accountId)
	ctx.IndentedJSON(http.StatusOK, data)
}
