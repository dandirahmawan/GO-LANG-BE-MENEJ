package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/input"
	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

func DeleteMember(ctx *gin.Context) {
	var input input.InputDeleteMember
	projectId := ctx.Param("projectId")
	userId := ctx.Param("userId")

	input.ProjectId = projectId
	input.UserId = userId

	data := services.DeleteMember(input)
	ctx.IndentedJSON(http.StatusOK, data)
}

func InviteMember(ctx *gin.Context) {
	var input input.InviteMember
	var sessionid = ctx.Request.Header.Get("sessionid")
	ctx.BindJSON(&input)

	datatoken := commons.GetDataSession(sessionid)
	data := services.InviteMember(input, datatoken.AccountId)
	ctx.IndentedJSON(http.StatusOK, data)
}

func Invitation(ctx *gin.Context) {
	id := ctx.Param("id")
	data := services.GetInvitation(id)
	ctx.IndentedJSON(http.StatusOK, data)
}
