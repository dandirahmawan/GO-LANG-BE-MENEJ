package model

type ViewLabelModule struct {
	ModuleId  int64  `json:"moduleId"`
	ProjectId int64  `json:"projectId"`
	Label     string `json:"label"`
	Color     string `json:"color"`
}

func (v ViewLabelModule) TableName() string {
	return "view_label_module"
}
