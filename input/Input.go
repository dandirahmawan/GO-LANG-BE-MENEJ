package input

type InputProject struct {
	ProjectName string `json:"projectName"`
}

type InputUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type InputDeleteMember struct {
	ProjectId string `json:"projectId"`
	UserId    string `json:"userId"`
}

type InviteMember struct {
	Email     string `json:"email"`
	ProjectId string `json:"projectId"`
}

type LabelDelete struct {
	ProjectId string `json:"projectId"`
	Label     string `json:"label"`
}
