package model

type ModelUserData struct {
	Org			string `json:"org"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	OldPassword string `json:"oldPassword"`
	Name		string `json:"name"`
	Mobile		string `json:"mobile"`
	Age			string `json:"age"`
	Salary		string `json:"salary"`
	Role	    string `json:"role"`
	Targets		string `json:"targets"`
}
