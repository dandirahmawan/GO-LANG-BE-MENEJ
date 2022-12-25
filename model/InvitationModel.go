package model

import (
	"time"
)

type InvitationModel struct {
	InvitationId    string    `json:"invitationId" gorm:"primaruKey"`
	UserId          string    `json:"userId"`
	InvitationEmail string    `json:"invitationEmail"`
	Status          string    `json:"status"`
	InvitationDate  time.Time `json:"invitationDate"`
	ProjectId       string    `json:"projectId"`
}

type ViewInvitationModel struct {
	InvitationId     string    `json:"invitationId" gorm:"primaruKey"`
	UserId           string    `json:"userId"`
	InvitationEmail  string    `json:"invitationEmail"`
	Status           string    `json:"status"`
	InvitationDate   time.Time `json:"invitationDate"`
	ProjectId        string    `json:"projectId"`
	ProjectName      string    `json:"projectName"`
	CreatedDate      time.Time `json:"createdDate"`
	UserNameInviting string    `json:"userNameInviting"`
	PicName          string    `json:"picName"`
}

func (m InvitationModel) TableName() string {
	return "invitation"
}

func (m ViewInvitationModel) TableName() string {
	return "view_invitation"
}

func (m InvitationModel) Save() InvitationModel {
	DB.Create(&m)
	return m
}

func (vm ViewInvitationModel) GetById() ViewInvitationModel {
	var data ViewInvitationModel
	DB.Where(ViewInvitationModel{InvitationId: vm.InvitationId}).Find(&data)
	return data
}
