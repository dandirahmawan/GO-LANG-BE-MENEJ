package model

type ViewAssignedModule struct {
	UserId   string `json:"userId"`
	ModuleId string `json:"moduleId"`
	UserName string `json:"userName"`
}

func (v ViewAssignedModule) TableName() string {
	return "view_assigned_module"
}
