package model

type ViewProjectTeam struct {
	ProjectId  string `json:"projectId" gorm:"primaryKey"`
	UserId     string `json:"userId" gorm:"primaryKey"`
	UserName   string `json:"userName"`
	EmailUser  string `json:"emailUser"`
	PicProfile string `json:"picProfile"`
	UserStatus string `json:"userStatus"`
}

func (v ViewProjectTeam) TableName() string {
	return "view_project_team"
}

func FindViewProjectTeamByProjectId(projectId string) []ViewProjectTeam {
	type Model ViewProjectTeam
	var Data []ViewProjectTeam
	DB.Where(&Model{ProjectId: projectId}).Find(&Data)

	return Data
}
