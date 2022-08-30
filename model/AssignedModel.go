package model

import "github.com/dandirahmawan/menej_api_go/config"

type AssignedModel struct {
	ModuleId string `json:"moduleId"`
	UserId   string `json:"userId"`
}

func (m AssignedModel) TableName() string {
	return "assigned_module"
}

func (m AssignedModel) Save() {
	db, _ := config.ConnectDB()
	db.Create(m)

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()
}

func DeleteAsignedByMoudlId(modulId string) {
	db, _ := config.ConnectDB()
	db.Delete(&AssignedModel{}, "module_id = ?", modulId)

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()
}
