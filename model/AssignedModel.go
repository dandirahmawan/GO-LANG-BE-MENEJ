package model

import "fmt"

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
	err := DB.Delete(&AssignedModel{}, "module_id = ?", modulId)
	if err != nil {
		fmt.Println(err)
	}
}
