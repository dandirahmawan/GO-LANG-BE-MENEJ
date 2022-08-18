package model

import "time"

type ViewModule struct {
	ModulId        int64                `json:"modulId" gorm:"primaryKey"`
	ProjectId      int64                `json:"projectId"`
	ModulName      string               `json:"modulName"`
	Pic            int64                `json:"pic"`
	CreatedBy      uint8                `json:"createdBy"`
	UpdatedBy      uint8                `json:"updatedBy"`
	CreatedDate    time.Time            `json:"createdDate"`
	UpdatedDate    time.Time            `json:"updatedDate"`
	EndDate        time.Time            `json:"endDate"`
	ModulStatus    string               `json:"modulStatus"`
	Description    string               `json:"description"`
	IsTrash        string               `json:"isTrash"`
	Percentage     string               `json:"percentage"`
	CountBugs      int64                `json:"countBugs"`
	CountDocFile   int64                `json:"countDocFile"`
	CountBugsClose int64                `json:"countBugsClose"`
	SectionId      int64                `json:"sectionId"`
	Label          []ViewLabelModule    `json:"label" gorm:"foreignKey:ModuleId;references:ModulId"`
	AssignTo       []ViewAssignedModule `json:"assignTo" gorm:"foreignKey:ModuleId;references:ModulId"`
}

func (v ViewModule) TableName() string {
	return "view_module"
}
