package model

import "github.com/dandirahmawan/menej_api_go/config"

type ViewLabelModule struct {
	ModuleId  string `json:"moduleId"`
	ProjectId string `json:"projectId"`
	Label     string `json:"label"`
	Color     string `json:"color"`
}

func (v ViewLabelModule) TableName() string {
	return "view_label_module"
}

func (v ViewLabelModule) FindByModulId() []ViewLabelModule {
	db, _ := config.ConnectDB()
	type M ViewLabelModule

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	var data []ViewLabelModule
	db.Where(&M{ModuleId: v.ModuleId}).Find(&data)
	return data
}
