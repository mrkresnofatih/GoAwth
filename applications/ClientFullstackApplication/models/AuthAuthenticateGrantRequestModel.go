package models

type AuthAuthenticateGrantRequestModel struct {
	GrantId           string `json:"grantId"`
	ApplicationId     string `json:"applicationId"`
	ApplicationSecret string `json:"applicationSecret"`
}
