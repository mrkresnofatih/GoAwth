package models

type DeveloperSignupRequestModel struct {
	DeveloperName string `json:"developerName" validate:"required,min=6,max=50"`
	Password      string `json:"password" validate:"required,containsany=abscdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=1234567890,containsany=!@#$%^&*(),min=6,max=20"`
}
