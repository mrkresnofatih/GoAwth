package models

type PlayerLoginResponseModel struct {
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}
