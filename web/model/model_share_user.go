package model

type ModelShareUserData struct {
	ShareOrg	string `json:"shareOrg"`
	Email		string `json:"email"`
	Access      int32  `json:"access"`
}
