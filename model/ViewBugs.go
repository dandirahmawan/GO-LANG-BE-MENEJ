package model

import (
	"time"
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
	type VB ViewBugs
	var data []ViewBugs
	DB.Where(&VB{ModulId: id}).Find(&data)
	return data
}
