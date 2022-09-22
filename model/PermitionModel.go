package model

type PermitionModel struct {
	ProjectId     string `json:"projectId" gorm:"primaryKey"`
	UserId        string `json:"userId" gorm:"primaryKey"`
	PermitionCode string `json:"permitionCode" gorm:"primaryKey"`
}

func (m PermitionModel) TableName() string {
	return "permition"
}

func (m PermitionModel) DeleteByProjectIdUserId() {
	var model PermitionModel
	DB.Where(PermitionModel{ProjectId: m.ProjectId, UserId: m.UserId}).Delete(&model)
}

func (m PermitionModel) Save() {
	DB.Save(&m)
}
