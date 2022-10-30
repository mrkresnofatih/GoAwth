package models

type PlayerSignupRequestModel struct {
	Username string `json:"username" validate:"required,min=6,max=30"`
	FullName string `json:"fullName" validate:"required,min=6,max=50"`
	ImageUrl string `json:"imageUrl" validate:"required"`
	Password string `json:"password" validate:"required,containsany=abscdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=1234567890,containsany=!@#$%^&*(),min=6,max=20"`
}
