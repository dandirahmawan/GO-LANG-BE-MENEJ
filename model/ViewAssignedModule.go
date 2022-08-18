package model

type ViewAssignedModule struct {
	UserId   int64  `json:"userId"`
	ModuleId int64  `json:"moduleId"`
	UserName string `json:"userName"`
}

func (v ViewAssignedModule) TableName() string {
	return "view_assigned_module"
}
