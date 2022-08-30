package model

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/config"
)

type BugsModel struct {
	BugsId     string    `json:"bugsId" gorm:"primaryKey"`
	ModulId    string    `json:"modulId"`
	ProjectId  string    `json:"projectId"`
	CreatedBy  string    `json:"createdBy"`
	CreateDate time.Time `json:"createDate"`
	IsDelete   string    `json:"isDelete"`
	BugStatus  string    `json:"bugStatus"`
	Note       string    `json:"note"`
}

func (m BugsModel) TableName() string {
	return "bugs"
}

func (m BugsModel) Save() {
	db, _ := config.ConnectDB()
	db.Save(m)

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()
}

func (m BugsModel) Update() {
	db, _ := config.ConnectDB()
	db.Updates(m)

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()
}
