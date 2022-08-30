package model

import (
	"github.com/dandirahmawan/menej_api_go/config"
)

type StatusModule struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	Color     string `json:"color"`
	Status    string `json:"status"`
}

func (m StatusModule) FindByProjectId() []StatusModule {
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	type M StatusModule
	var data []StatusModule
	db.Table("status_module").Where(&M{ProjectId: m.ProjectId}).Find(&data)
	return data
}
