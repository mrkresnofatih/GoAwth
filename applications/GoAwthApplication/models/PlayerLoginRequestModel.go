package models

type PlayerLoginRequestModel struct {
	Username string `json:"username" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}
