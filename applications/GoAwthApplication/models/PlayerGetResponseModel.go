package models

type PlayerGetResponseModel struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	ImageUrl string `json:"imageUrl"`
}
