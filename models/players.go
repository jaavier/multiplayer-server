package models

type Player struct {
	UserId 		string `json:"userId" required:"true"`
	Status 		string `json:"status"`
	LastLogin string `json:"lastLogin"`
}
