package model

import (
	"github.com/dandirahmawan/menej_api_go/config"
)

// var DB, _ = config.ConnectDB()

type LabelModule struct {
	Label     string `json:"label"`
	ModuleId  string `json:"moduleId" gorm:"primaryKey"`
	ProjectId string `json:"projectId" gorm:"primaryKey"`
}

func (m LabelModule) TableName() string {
	return "label_module"
}

func DeleteLabelModuleByModuleId(modulId string) {
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	db.Delete(LabelModule{}, "module_id = ?", modulId)
}

func (m LabelModule) Save() {
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	db.Create(&m)
}
