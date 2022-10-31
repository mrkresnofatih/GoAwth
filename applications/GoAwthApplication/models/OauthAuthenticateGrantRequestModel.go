package models

type OauthAuthenticateGrantRequestModel struct {
	GrantId           string `json:"grantId" validate:"required,max=100"`
	ApplicationId     string `json:"applicationId" validate:"required,max=100"`
	ApplicationSecret string `json:"applicationSecret" validate:"required,max=100"`
}
