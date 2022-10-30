package models

type DeveloperLoginResponseModel struct {
	DeveloperName string `json:"developerName"`
	AccessToken   string `json:"accessToken"`
}
