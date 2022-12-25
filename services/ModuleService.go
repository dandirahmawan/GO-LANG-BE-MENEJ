package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dandirahmawan/menej_api_go/commons"
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
	DataSection      interface{} `json:"sectionModules"`
	AssignedModules  interface{} `json:"assignedModules"`
}

type ParamModuleInput struct {
	Checklist []model.BugsModel `json:"checklist"`
	Module    model.ModulInput  `json:"module"`
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

func GetDataModuleDetail(id string, userid string) interface{} {
	mi := id
	userId := userid
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
	dataSection := model.FindSectionOnlyByProjectId(projectId)

	var resp ResultModuleDetail
	resp.Bugs = dataBugs
	resp.DataModule = dataModule
	resp.DataStatus = dataStatus
	resp.LabelModule = dataLabels
	resp.PermitionProject = dataPermition
	resp.AssignedModules = dataAssigned
	resp.DocumentFile = dataDocFile
	resp.DataSection = dataSection
	return resp
}

func SaveModule(ctx *gin.Context) interface{} {
	var param ParamModuleInput
	err := ctx.BindJSON(&param)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	moduleInput := param.Module
	checklist := param.Checklist
	assigned := moduleInput.Assigned

	var mod model.ModulModel
	var data model.ModulModel

	/*get data session*/
	sessionid := ctx.Request.Header.Get("sessionid")
	dataSession := commons.GetDataSession(sessionid)

	// fmt.Println(data)
	id := ctx.Param("id")
	var mId string
	if id == "" {
		mId = commons.GeneratdUUID(20)
	} else {
		mId = id
		data = mod.FindById()
	}

	// fmt.Println(data)
	data.ModulId = mId
	data.ModulName = moduleInput.ModuleName
	data.ProjectId = moduleInput.ProjectId
	data.Description = moduleInput.Desc
	data.ModulStatus = moduleInput.Status
	data.UpdatedDate = time.Now()
	data.SectionId = moduleInput.Section
	data.UpdatedBy = dataSession.AccountId
	data.IsTrash = "N"
	// data.OldSectionId = moduleInput.OldSectionId

	if id == "" {
		data.CreatedBy = dataSession.AccountId
		data.CreatedDate = time.Now()
	}

	/*set end or due date module*/
	// fmt.Println("input date :", moduleInput.Date)
	endDate, err := time.Parse("2006-01-02", moduleInput.Date)
	data.EndDate = endDate

	if err != nil {
		fmt.Println(err)
	}

	if id == "" {
		data.SaveModule()
	} else {
		data.Update()
	}

	/*set data label module*/
	setDataLabel(moduleInput.LabelModule, data)

	/*set data assigned*/
	setDataAssigned(assigned, data)

	/*set data checklist*/
	setDataChecklist(checklist, data)

	var resp interface{} = GetDataModuleDetail(mId, dataSession.AccountId)
	return resp
}

func setDataAssigned(assigned string, module model.ModulModel) {
	type JsonAssigned struct {
		UserId   string `json:"userId"`
		ModuleId string `json:"moduleId"`
	}

	var arr []JsonAssigned
	err := json.Unmarshal([]byte(assigned), &arr)

	if err != nil {
		fmt.Println(err)
	}

	if len(arr) > 0 {
		model.DeleteAsignedByMoudlId(module.ModulId)
		for i := 0; i < len(arr); i++ {
			item := arr[i]

			var am model.AssignedModel
			am.ModuleId = module.ModulId
			am.UserId = item.UserId
			am.Save()
		}
	}
}

func setDataLabel(labelModule string, module model.ModulModel) {
	var arrLabel []model.LabelModule
	json.Unmarshal([]byte(labelModule), &arrLabel)
	if len(arrLabel) > 0 {
		/*delete old data*/
		model.DeleteLabelModuleByModuleId(module.ModulId)

		for i := 0; i < len(arrLabel); i++ {
			item := arrLabel[i]
			item.ModuleId = module.ModulId
			item.ProjectId = module.ProjectId
			item.Save()
		}
	}
}

func setDataChecklist(checklist []model.BugsModel, module model.ModulModel) {
	for i := 0; i < len(checklist); i++ {
		item := checklist[i]
		bugsId := item.BugsId

		if bugsId == "-1" {
			bi := commons.GeneratdUUID(20)
			item.BugsId = bi
			item.ModulId = module.ModulId
			item.ProjectId = module.ProjectId
			item.Save()
		} else {
			item.Update()
		}
	}
}

func DeleteModuleById(moduleid string) interface{} {
	var mod model.ModulModel
	mod.ModulId = moduleid
	i := mod.DeleteById()

	type resp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	var rsp resp
	if i == 0 {
		rsp.Message = "Somethin wrong when delete module"
		rsp.Success = false
	} else {
		rsp.Message = "Delete module successfully"
		rsp.Success = true
	}

	return rsp
}
