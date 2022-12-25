package model

import (
	"fmt"
	"time"
)

type ProjectTeamModel struct {
	UserId      string    `json:"userId"`
	ProjectId   string    `json:"projectId"`
	CreatedDate time.Time `json:"createdDate"`
	UserStatus  string    `json:"userStatus"`
}

func (m ProjectTeamModel) TableName() string {
	return "project_team"
}

func (m ProjectTeamModel) Delete() int {
	fmt.Println(m)
	err := DB.Where(ProjectTeamModel{ProjectId: m.ProjectId, UserId: m.UserId}).Updates(&m)
	if err != nil {
		return 0
	}

	return 1
}
