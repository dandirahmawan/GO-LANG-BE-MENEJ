package model

type ViewLabelModule struct {
	ModuleId  string `json:"moduleId"`
	ProjectId string `json:"projectId"`
	Label     string `json:"label"`
	Color     string `json:"color"`
}

func (v ViewLabelModule) TableName() string {
	return "view_label_module"
}

func (v ViewLabelModule) FindByModulId() []ViewLabelModule {
	type M ViewLabelModule

	var data []ViewLabelModule
	DB.Where(&M{ModuleId: v.ModuleId}).Find(&data)
	return data
}
