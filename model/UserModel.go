package model

import (
	"time"
)

type UserModel struct {
	UserId           string    `json:"userId" gorm:"primaryKey"`
	EmailUser        string    `json:"emailUser"`
	UserName         string    `json:"userName"`
	UserPassword     string    `json:"userPassword"`
	UserType         string    `json:"userType"`
	PicProfile       string    `json:"picProfile"`
	SessionId        string    `json:"sessionId"`
	Date             time.Time `json:"date"`
	PicProfileDetail string    `json:"picProfileDetail"`
}

type TableName interface {
	TableName()
}

func (u UserModel) TableName() string {
	return "user"
}
