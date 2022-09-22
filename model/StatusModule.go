package model

type StatusModule struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	Color     string `json:"color"`
	Status    string `json:"status"`
}

func (m StatusModule) FindByProjectId() []StatusModule {
	type M StatusModule
	var data []StatusModule
	DB.Table("status_module").Where(&M{ProjectId: m.ProjectId}).Find(&data)
	return data
}
