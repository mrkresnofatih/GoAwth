package models

type DeveloperAppGrantCreateResponseModel struct {
	DeveloperAppGrantId string
	Username            string
	ApplicationId       string
	ExpiresAt           string
	Scope               string
}
