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
	IsConfirmed      int16     `json:"isConfirmed"`
}

type TableName interface {
	TableName()
}

func (u UserModel) TableName() string {
	return "user"
}

func (u UserModel) Save() UserModel {
	DB.Create(&u)
	return u
}

func (u UserModel) FindByEmail() []UserModel {
	var data []UserModel
	DB.Where(UserModel{EmailUser: u.EmailUser}).Find(&data)
	return data
}

func (u UserModel) FindById() UserModel {
	var data UserModel
	DB.Where(UserModel{UserId: u.UserId}).Find(&data)
	return data
}
