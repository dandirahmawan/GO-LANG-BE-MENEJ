package model

import "fmt"

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

func FindLabelModuleByModuleId(modulId string) []LabelModule {
	var data []LabelModule
	err := DB.Where(LabelModule{ModuleId: modulId}).Find(&data)

	if err != nil {
		fmt.Println(err)
	}

	return data
}

func DeleteLabeModulelByModuleId(modulId string) {
	DB.Delete(&LabelModule{ModuleId: modulId})
}

func (m LabelModule) DeleteByProjectIdAndLabel() {
	err := DB.Delete(&LabelModule{ProjectId: m.ProjectId, Label: m.Label})
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
}

func (m LabelModule) Save() {
	DB.Create(&m)
}
