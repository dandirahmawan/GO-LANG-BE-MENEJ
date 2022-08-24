package model

import "github.com/dandirahmawan/menej_api_go/config"

type ViewAssignedModule struct {
	UserId    string `json:"userId"`
	ProjectId string `json:"projectId"`
	ModuleId  string `json:"moduleId"`
	UserName  string `json:"userName"`
	EmailUser string `json:"emailUser"`
}

func (v ViewAssignedModule) TableName() string {
	return "view_assigned_module"
}

func (v ViewAssignedModule) FindByModulId() []ViewAssignedModule {
	type M ViewAssignedModule
	db, _ := config.ConnectDB()

	var data []ViewAssignedModule
	db.Where(M{ModuleId: v.ModuleId}).Find(&data)
	return data
}
