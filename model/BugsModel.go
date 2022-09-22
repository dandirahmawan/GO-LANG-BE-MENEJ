package model

import (
	"time"
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
	DB.Save(m)
}

func (m BugsModel) Update() {
	DB.Updates(m)
}
