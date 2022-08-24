package services

import (
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

type ResultModuleDetail struct {
	Bugs             interface{} `json:"bugs"`
	PermitionProject interface{} `json:"permitionProject"`
	DataModule       interface{} `json:"dataModule"`
	DataStatus       interface{} `json:"dataStatus"`
	LabelModule      interface{} `json:"labelModules"`
	DocumentFile     interface{} `json:"documentFile"`
	AssignedModules  interface{} `json:"assignedModules"`
}

func GetModuleByProjecId(projectid string) []model.SectionModel {
	var data []model.SectionModel
	data = model.FindSectionByProjectId(projectid)
	return data
}

func GetDataStatus(ctx *gin.Context) interface{} {
	id := ctx.Param("id")
	var sm model.StatusModule
	sm.ProjectId = id
	return sm.FindByProjectId()
}

func GetDataModuleDetail(ctx *gin.Context) interface{} {
	mi := ctx.Param("id")
	userId := ctx.Request.Header.Get("userid")
	var bugs model.ViewBugs
	var module model.ViewModule
	var status model.StatusModule
	var label model.ViewLabelModule
	var Assign model.ViewAssignedModule
	var docFile model.ViewDocumentFile

	bugs.ModulId = mi
	module.ModulId = mi
	label.ModuleId = mi
	Assign.ModuleId = mi
	docFile.ModulId = mi

	dataBugs := bugs.FindByModulId()
	dataModule := module.FindById()
	dataLabels := label.FindByModulId()
	dataAssigned := Assign.FindByModulId()
	dataDocFile := docFile.FindByModulId()

	/*get project id as param*/
	projectId := dataModule.ProjectId
	status.ProjectId = projectId
	dataStatus := status.FindByProjectId()
	dataPermition := model.FindDataPermitionByProjectIdAndUserId(projectId, userId)

	var resp ResultModuleDetail
	resp.Bugs = dataBugs
	resp.DataModule = dataModule
	resp.DataStatus = dataStatus
	resp.LabelModule = dataLabels
	resp.PermitionProject = dataPermition
	resp.AssignedModules = dataAssigned
	resp.DocumentFile = dataDocFile
	return resp
}
