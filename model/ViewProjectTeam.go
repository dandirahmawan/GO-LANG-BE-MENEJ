package model

import "github.com/dandirahmawan/menej_api_go/config"

type ViewProjectTeam struct {
	ProjectId  int64  `json:"projectId" gorm:"primaryKey"`
	UserId     int64  `json:"userId" gorm:"primaryKey"`
	UserName   string `json:"userName"`
	EmailUser  string `json:"emailUser"`
	PicProfile string `json:"picProfile"`
	UserStatus string `json:"userStatus"`
}

func (v ViewProjectTeam) TableName() string {
	return "view_project_team"
}

func FindViewProjectTeamByProjectId(projectId int64) []ViewProjectTeam {
	db, _ := config.ConnectDB()

	type Model ViewProjectTeam
	var Data []ViewProjectTeam
	db.Where(&Model{ProjectId: projectId}).Find(&Data)

	return Data
}
