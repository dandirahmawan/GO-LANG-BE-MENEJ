package services

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/input"
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/dandirahmawan/menej_api_go/response_base"
	"github.com/gin-gonic/gin"
)

type respMsg response_base.RespMsg
type respData response_base.Resp

func GetProjectTeamByProjectId(ctx *gin.Context) interface{} {
	var projectID string = ctx.Param("id")
	data := model.FindViewProjectTeamByProjectId(projectID)
	return data
}

func InviteMember(param input.InviteMember, userid string) respMsg {
	uuid := commons.GeneratdUUID(20)
	var mod model.InvitationModel
	mod.InvitationId = uuid
	mod.ProjectId = param.ProjectId
	mod.InvitationEmail = param.Email
	mod.Status = "PENDING"
	mod.InvitationDate = time.Now()
	mod.UserId = userid
	mod.Save()

	/*get data view invitation*/
	var modv model.ViewInvitationModel
	modv.InvitationId = mod.InvitationId
	dataView := modv.GetById()
	go SendEmailInvitation(dataView)

	var msg respMsg
	msg.Success = true
	msg.Message = "Invitation is successfuly"
	return msg
}

func GetInvitation(id string) respData {
	var data model.ViewInvitationModel
	data.InvitationId = id
	data = data.GetById()

	/*get data team*/
	projectId := data.ProjectId
	dataTeam := model.FindViewProjectTeamByProjectId(projectId)

	type dataInvitation struct {
		DataProjectTeamList interface{} `json:"dataProjectTeamList"`
		ViewInvitation      interface{} `json:"viewInvitation"`
	}

	var dataInv dataInvitation
	dataInv.ViewInvitation = data
	dataInv.DataProjectTeamList = dataTeam

	var msg respData
	msg.Success = true
	msg.Message = "Invitation is successfuly"
	msg.Data = dataInv
	return msg
}

func DeleteMember(param input.InputDeleteMember) response_base.RespMsg {
	var mod model.ProjectTeamModel
	mod.ProjectId = param.ProjectId
	mod.UserId = param.UserId
	mod.UserStatus = "N"
	exc := mod.Delete()

	if exc == 1 {
		return response_base.ResponseMsg(true, "Delete Member Successfully")
	} else {
		return response_base.ResponseMsg(false, "Delete Member Failed")
	}
}
