package input

type InputProject struct {
	ProjectName string `json:"projectName"`
}

type InputUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
