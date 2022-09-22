package model

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
	var data []ViewAssignedModule
	DB.Where(M{ModuleId: v.ModuleId}).Find(&data)
	return data
}
