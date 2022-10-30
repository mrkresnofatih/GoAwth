package models

type DeveloperLoginRequestModel struct {
	DeveloperName string `json:"developerName" validate:"required,min=6,max=30"`
	Password      string `json:"password" validate:"required,containsany=abscdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=1234567890,containsany=!@#$%^&*(),min=6,max=20"`
}
