package services

import (
	"fmt"
	"time"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/input"
	"github.com/dandirahmawan/menej_api_go/model"
)

func InsertUser(param input.InputUser) model.UserModel {
	var userModel model.UserModel
	userModel.EmailUser = param.Email
	userModel.UserName = param.Name
	userModel.UserPassword = param.Password
	userModel.Date = time.Now()
	userModel.IsConfirmed = 0
	userModel.UserId = commons.GeneratdUUID(20)
	exc := userModel.Save()
	return exc
}

func RegisterUser(param input.InputUser) interface{} {
	var userModel model.UserModel
	userModel.EmailUser = param.Email
	data := userModel.FindByEmail()

	type resp struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	var response resp
	if (len(data)) > 0 {
		response.Success = false
		response.Message = "Email Is Not Available"
	} else {
		data := InsertUser(param)

		/*Init data project*/
		var inputProject input.InputProject
		inputProject.ProjectName = "My First Project"

		project := SaveProject(inputProject, data.UserId)
		var section model.SectionModel
		section.Id = commons.GeneratdUUID(20)
		section.Section = "TO DO LIST"
		section.CreatedDate = time.Now()
		section.ProjectId = project.ProjectId
		SaveSection(section)

		response.Success = true
		response.Message = "Register Is Successfully"
		response.Data = data
	}
	return response
}

func FinDByEmailAndPassword(email string, pass string) interface{} {

	var data []model.UserModel
	type Mod model.UserModel

	DB.Where(&Mod{EmailUser: email, UserPassword: pass}).Find(&data)

	count := len(data)
	if count <= 0 {
		type resp struct {
			Code    int16  `json:"code"`
			Message string `json:"message"`
		}

		var rsp resp
		var rspList []resp
		rsp.Code = 201
		rsp.Message = "User not registered"

		rspList = append(rspList, rsp)
		return rspList
	}

	/*response success login*/
	type s struct {
		Code      int16  `json:"code"`
		SessionId string `json:"sessionId"`
		UserId    string `json:"userId"`
	}

	accountId := data[0].UserId
	token := commons.GeneratdUUID(30)
	fmt.Println("token ", token)

	/*sava data session*/
	SaveSessionLogin(accountId, token)

	var r s
	r.Code = 200
	r.SessionId = token
	r.UserId = accountId

	/*arra respones login*/
	var arr []s
	arr = append(arr, r)

	return arr
}

func FindAllUser() []model.UserModel {

	var data []model.UserModel

	DB.Find(&data)
	return data
}

func SaveSessionLogin(accountId string, token string) {
	now := time.Now()
	exp := now.AddDate(1, 0, 0)

	var sm model.SessionModel
	sm.AccountId = accountId
	sm.ExpiredDate = exp
	sm.Id = token

	model.SaveSession(sm)
}
