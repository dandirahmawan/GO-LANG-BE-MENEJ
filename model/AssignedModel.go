package model

type AssignedModel struct {
	ModuleId string `json:"moduleId"`
	UserId   string `json:"userId"`
}

func (m AssignedModel) TableName() string {
	return "assigned_module"
}

func (m AssignedModel) Save() {
	DB.Create(m)
}

func DeleteAsignedByMoudlId(modulId string) {
	DB.Delete(&AssignedModel{}, "module_id = ?", modulId)
}
