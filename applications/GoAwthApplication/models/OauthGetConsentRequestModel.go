package models

type OauthGetConsentRequestModel struct {
	DeveloperApplicationId string `json:"developerApplicationId" validate:"required,max=50"`
	Scope                  string `json:"scope" validate:"required,max=200"`
	GrantType              string `json:"grantType" validate:"required"`
	PlayerUsername         string `json:"playerUsername" validate:"required"`
}
