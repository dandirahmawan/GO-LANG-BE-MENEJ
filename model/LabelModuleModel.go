package model

// var DB, _ = config.ConnectDB()

type LabelModule struct {
	Label     string `json:"label"`
	ModuleId  string `json:"moduleId" gorm:"primaryKey"`
	ProjectId string `json:"projectId" gorm:"primaryKey"`
}

func (m LabelModule) TableName() string {
	return "label_module"
}

func DeleteLabelModuleByModuleId(modulId string) {
	DB.Delete(LabelModule{}, "module_id = ?", modulId)
}

func (m LabelModule) Save() {
	DB.Create(&m)
}
