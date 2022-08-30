package model

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/config"
)

type ViewBugs struct {
	BugsId     string    `json:"bugsId"`
	BugStatus  string    `json:"bugStatus"`
	CreateDate time.Time `json:"createDate"`
	IsDelete   string    `json:"isDelete"`
	ModulId    string    `json:"modulId"`
	Note       string    `json:"note"`
}

func (m ViewBugs) TableName() string {
	return "view_bugs"
}

func (m ViewBugs) FindByModulId() []ViewBugs {
	id := m.ModulId
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	type VB ViewBugs
	var data []ViewBugs
	db.Where(&VB{ModulId: id}).Find(&data)
	return data
}
